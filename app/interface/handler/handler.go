package _interface

import (
	"encoding/json"
	"example.com/application/auth"
	"example.com/application/service"
	"example.com/interface/request"
	"example.com/interface/response"
	"github.com/uptrace/bunrouter"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: *userService}
}

// UserCreateHandle creates a new user.
//
//	@Summary		Create a new user
//	@Description	Create a new user with the given name.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			name	body		request.UserCreateRequest	true	"User info"
//	@Success		200		{object}	response.UserCreateResponse
//	@Router			/user/create [post]
func (u *UserHandler) UserCreateHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var requestData request.UserCreateRequest
		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return err
		}

		ctx := req.Context()
		authToken, err := u.userService.Add(ctx, requestData.Name)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return err
		}

		responseData := &response.UserCreateResponse{Token: authToken}
		responseBytes, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return nil
	}
}

// UserGetHandle retrieves the information of a user.
//
//	@Summary		Get user information
//	@Description	Get the information of a user by their ID.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.UserGetResponse
//	@Router			/user/get [post]
func (u *UserHandler) UserGetHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		id := auth.GetUserIDFromContext(ctx)

		// Retrieve user by auth token
		user, err := u.userService.GetUserByUserId(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		// Prepare the response using UserGetResponse struct
		responseData := &response.UserGetResponse{
			Id:        user.Id,
			Name:      user.Name,
			HighScore: user.HighScore,
		}

		respBytes, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)

		return nil
	}
}

// ScoreUpdateHandle updates the score of a user.
//
//	@Summary		Update user score
//	@Description	Update the score of a user if the new score is higher than the current score.
//	@Tags			Scores
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{string}	string	"ScoreUpdateHandle triggered"
//	@Router			/user/score [post]
func (u *UserHandler) ScoreUpdateHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		w.Write([]byte("ScoreUpdateHandle triggered"))
		var requestData request.UserScoreUpdateRequest
		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return err
		}
		ctx := req.Context()
		id := auth.GetUserIDFromContext(ctx)

		user, err := u.userService.GetUserByUserId(ctx, id)
		if err != nil {
			http.Error(w, "Failed to Score Update user", http.StatusInternalServerError)
			return err
		}

		score, _ := strconv.Atoi(requestData.Score)

		if user.HighScore < score {
			user.HighScore = score
			_ = u.userService.UpdateUser(ctx, user)
		}

		return nil
	}
}

// UserRankingGetHandle retrieves the rankings of all users.
//
//	@Summary		Get user rankings
//	@Description	Get the rankings of all users based on their high scores.
//	@Tags			Rankings
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	response.UserRankingResponse
//	@Router			/users/get [get]
func (u *UserHandler) UserRankingGetHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		// UserServiceからランキングを取得
		userRankings, err := u.userService.GetUserRanking(req.Context())
		if err != nil {
			http.Error(w, "Failed to get user rankings", http.StatusInternalServerError)
			return err
		}

		// UserRankingからUserRankingResponseに変換
		var responseSlice []response.UserRankingResponse
		for _, ranking := range userRankings {
			r := response.UserRankingResponse{
				Name:      ranking.Name,
				HighScore: ranking.HighScore,
			}
			responseSlice = append(responseSlice, r)
		}

		// ヘッダーを設定してJSONとしてレスポンスを返すことを示す
		w.Header().Set("Content-Type", "application/json")

		// ランキングをJSONとしてエンコードしてレスポンスに書き込む
		if err := json.NewEncoder(w).Encode(responseSlice); err != nil {
			http.Error(w, "Failed to encode user rankings", http.StatusInternalServerError)
			return err
		}

		return nil
	}
}

// DestroyHandle deletes a user.
//
//	@Summary		Delete a user
//	@Description	Delete a user by their ID.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	map[string]string
//	@Router			/user/destroy [post]
func (u *UserHandler) DestroyHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := req.Context()
		id := auth.GetUserIDFromContext(ctx)

		// Delete the user using userService
		if _, err := u.userService.Delete(ctx, id); err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "User successfully deleted"}`))
		return nil
	}
}
