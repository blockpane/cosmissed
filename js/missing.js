async function whosMissing() {

    let wsProto = "ws://"
    if (location.protocol === "https:") {
        wsProto = "wss://"
    }
    const socket = new WebSocket(wsProto+location.host+'/missed/ws');
    socket.addEventListener('message', function (event) {
        const missing = document.getElementById('missing');
        missing.innerHTML = "";
        let upd = JSON.parse(event.data);
        let monikers = [];
        for (const [key, value] of Object.entries(upd.missing)) {
            monikers.push(key)
        }
        if (monikers.length > 0) {
            monikers.sort(function(a, b) {
                let nameA = a.name.toUpperCase();
                let nameB = b.name.toUpperCase();
                if (nameA < nameB) {return -1;}
                if (nameA > nameB) {return 1;}
                return 0;
            });
            monikers.forEach((moniker) => {
                let li = document.createElement("li")
                li.appendChild(document.createTextNode(moniker));
                missing.appendChild(li);
                console.log(moniker+" is missing signatures");
            });
        }
        document.getElementById('headblock').innerHTML = upd.block_num;
        document.getElementById('seconds').innerHTML = upd.delta_sec;
    });
}
