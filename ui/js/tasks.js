const API_URL = "http://localhost:8080/tasks";

function getToken() {
    return localStorage.getItem("token");
}

function logout() {
    localStorage.removeItem("token");
    window.location.href = "index.html";
}

// 🔐 redirect if not logged in
if (!getToken()) {
    window.location.href = "index.html";
}

// 🚀 LOAD TASKS AUTOMATICALLY
async function loadTasks() {
    const res = await fetch(API_URL, {
        headers: {
            "Authorization": "Bearer " + getToken()
        }
    });

    if (!res.ok) {
        console.error("Failed:", res.status);
        return;
    }

    const result = await res.json();
    const tasks = result.data || [];

    const list = document.getElementById("taskList");
    list.innerHTML = "";

    tasks.forEach(task => {
        const li = document.createElement("li");

        li.innerHTML = `
            <span style="${task.completed ? 'text-decoration:line-through;' : ''}">
                ${task.title}
            </span>

            <div>
                <button onclick="toggleTask(${task.id}, ${task.completed})">
                    ${task.completed ? "Undo" : "Done"}
                </button>

                <button onclick="deleteTask(${task.id})">
                    Delete
                </button>
            </div>
        `;

        list.appendChild(li);
    });
}

// ➕ ADD
async function addTask() {
    const input = document.getElementById("taskInput");

    if (!input.value.trim()) return;

    await fetch(API_URL, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + getToken()
        },
        body: JSON.stringify({
            title: input.value,
            completed: false
        })
    });

    input.value = "";
    loadTasks(); // 🔥 auto refresh
}

// ❌ DELETE
async function deleteTask(id) {
    await fetch(`${API_URL}/${id}`, {
        method: "DELETE",
        headers: {
            "Authorization": "Bearer " + getToken()
        }
    });

    loadTasks();
}

// 🔁 UPDATE
async function toggleTask(id, completed) {
    await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + getToken()
        },
        body: JSON.stringify({
            completed: !completed
        })
    });

    loadTasks();
}

// 🔥 LOAD ON PAGE OPEN
loadTasks();