const URI = "/api/"

async function checkLogin() {
    let key = localStorage.getItem("key");
    if (key == null) {
        return Promise.reject("No key found in local storage")
    } else {
        const response = await fetch(URI, {
            headers: {
                "key": key
            }
        });
        if (!response.ok) {
            return Promise.reject("Login Failed.");
        } else {
            return response.json();
        }
    }
}

async function getPlayers() {
    const response = await fetch(URI + "player", {
        headers: {
            "key": localStorage.getItem("key"),
        }
    });
    if (!response.ok) {
        return Promise.reject(response.statusText);
    } else {
        return response.json();
    }
}

async function getPlayers(id) {
    const response = await fetch(URI + "player/" + id, {
        "key": localStorage.getItem("key"),
    });
    if (!response.ok) {
        return Promise.reject(response.statusText);
    } else {
        return response.json();
    }
}