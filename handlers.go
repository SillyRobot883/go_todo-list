package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// GET all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, title, description, complete FROM tasks")
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// to show the tasks in a JSON format
	var tasks []Tasks
	for rows.Next() {
		var task Tasks
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.isCompleted); err != nil {
			http.Error(w, "Error reading tasks", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GET specific task by id
func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task Tasks
	query := DB.QueryRow("SELECT id, title, description, complete FROM tasks WHERE id = ?", id)
	err := query.Scan(&task.ID, &task.Title, &task.Description, &task.isCompleted)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching task", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Tasks
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec("INSERT INTO tasks (title, description, complete) VALUES (?, ?, ?)", task.Title, task.Description, task.isCompleted)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH update a task
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task Tasks
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec("UPDATE tasks SET title = ?, description = ?, complete = ? WHERE id = ?", task.Title, task.Description, task.isCompleted, id)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
