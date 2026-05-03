const API_URL = "http://localhost:8080/tasks";

function logout() {
    localStorage.removeItem("token");
    window.location.href = "login.html";
}

async function loadTasks() {
    const res = await fetch(API_URL);
    const tasks = await res.json();

    const list = document.getElementById("taskList");
    const count = document.getElementById("taskCount");

    list.innerHTML = "";
    count.textContent = `${tasks.length} Tasks`;

    tasks.forEach(task => {
        const li = document.createElement("li");

        if (task.completed) {
            li.classList.add("completed");
        }

        li.innerHTML = `
            <span>${task.title}</span>
            <div class="actions">
                <button class="complete-btn" onclick="toggleTask(${task.id}, ${task.completed})">
                    ${task.completed ? "Undo" : "Done"}
                </button>
                <button class="delete-btn" onclick="deleteTask(${task.id})">
                    Delete
                </button>
            </div>
        `;

        list.appendChild(li);
    });
}

async function addTask() {
    const input = document.getElementById("taskInput");

    if (!input.value.trim()) return;

    await fetch(API_URL, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            title: input.value,
            completed: false
        })
    });

    input.value = "";
    loadTasks();
}

async function deleteTask(id) {
    await fetch(`${API_URL}/${id}`, { method: "DELETE" });
    loadTasks();
}

async function toggleTask(id, completed) {
    await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            completed: !completed
        })
    });
    loadTasks();
}

loadTasks();