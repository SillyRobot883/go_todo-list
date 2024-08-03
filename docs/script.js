const apiUrl = 'http://localhost:8085/api/v1/task';

document.addEventListener('DOMContentLoaded', () => {
    fetchTasks();
});

function fetchTasks() {
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            const taskList = document.getElementById('task-list');
            taskList.innerHTML = '';
            data.forEach(task => {
                const taskItem = document.createElement('div');
                taskItem.className = 'task-item';
                if (task.isCompleted) taskItem.classList.add('completed');

                taskItem.innerHTML = `
                    <div>
                        <strong>${task.title}</strong><br>
                        ${task.description}
                    </div>
                    <div>
                        <button onclick="toggleTask(${task.id}, ${task.isCompleted})">${task.isCompleted ? 'Mark as Incomplete' : 'Mark as Complete'}</button>
                        <button onclick="deleteTask(${task.id})">Delete</button>
                    </div>
                `;
                taskList.appendChild(taskItem);
            });
        })
        .catch(error => console.error('Error fetching tasks:', error));
}

function addTask() {
    const title = document.getElementById('task-title').value;
    const description = document.getElementById('task-description').value;

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title, description, isCompleted: false })
    })
        .then(response => {
            if (response.ok) {
                fetchTasks();
                document.getElementById('task-title').value = '';
                document.getElementById('task-description').value = '';
            } else {
                console.error('Error adding task:', response.statusText);
            }
        })
        .catch(error => console.error('Error adding task:', error));
}

function toggleTask(id, isCompleted) {
    fetch(`${apiUrl}/${id}`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ isCompleted: !isCompleted })
    })
        .then(response => {
            if (response.ok) {
                fetchTasks();
            } else {
                console.error('Error toggling task:', response.statusText);
            }
        })
        .catch(error => console.error('Error toggling task:', error));
}

function deleteTask(id) {
    fetch(`${apiUrl}/${id}`, {
        method: 'DELETE'
    })
        .then(response => {
            if (response.ok) {
                fetchTasks();
            } else {
                console.error('Error deleting task:', response.statusText);
            }
        })
        .catch(error => console.error('Error deleting task:', error));
}
