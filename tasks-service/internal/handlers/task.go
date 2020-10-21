package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/task-management-microservices/tasks-service/internal/domain"
	"net/http"
	"strconv"
)

type createTaskResponse struct {
	Id int `json:"id"`
}

func (h *Handler) createTask(c *gin.Context) {
	var input domain.Task
	if err := c.BindJSON(&input); err != nil {
		responseAndLogError(c, http.StatusBadRequest, "createTask", "invalid input body: "+err.Error())
		return
	}

	input.UserId = c.GetInt(userIdCtx)

	id, err := h.taskService.CreateTask(input)
	if err != nil {
		responseAndLogError(c, http.StatusInternalServerError, "createTask", err.Error())
		return
	}

	c.JSON(http.StatusOK, createTaskResponse{id})
}

type getAllTasksResponse struct {
	Tasks []domain.Task `json:"tasks"`
}

func (h *Handler) getAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAll()
	if err != nil {
		responseAndLogError(c, http.StatusInternalServerError, "getAllTasks", err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTasksResponse{tasks})
}

func (h *Handler) getTaskById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responseAndLogError(c, http.StatusBadRequest, "getTaskById", "invalid id url parameter")
		return
	}

	task, err := h.taskService.GetById(id)
	if err != nil {
		responseAndLogError(c, http.StatusInternalServerError, "getTaskById", err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
