function drawValidators(vals) {
    const infoCircle = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-info-circle" viewBox="0 0 16 16">
  <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
  <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
</svg>`
    document.getElementById("validators").innerHTML = null
    vals.forEach ((v) => {
        let row = document.createElement("tr")
        let name = row.insertCell(-1)
        name.className = "text-start"
        name.innerHTML = `<button class="btn-light">${infoCircle}</button> &nbsp;` + v["name"]
        row.insertCell(-1).innerHTML = v["stake"]
        row.insertCell(-1).innerHTML = v["self_stake"]
        row.insertCell(-1).innerHTML = v["missed"]
        row.insertCell(-1).innerHTML = v["missed_pct"]+" %"
        row.insertCell(-1).innerHTML = v["unclaimed_rew"]
        row.insertCell(-1).innerHTML = v["unclaimed_com"]
        row.insertCell(-1).innerHTML = v["votes"]
        document.getElementById("validators").appendChild(row)
    })
}

function filterTable(minStake = 0, maxStake = 0, minSelf = 0, maxSelf = 0, missed = 0, gov = false) {
    let tbl = document.getElementById("sortTable")
    for (let i = 0; i < tbl.rows.length; i++) {
        const stake = parseInt(tbl.rows[i].cells[1].innerHTML)
        if ((stake < minStake && minStake !== 0) || (stake >= maxStake && maxStake !== 0)) {
            tbl.deleteRow(i)
            --i
            continue
        }
        const selfStake = parseInt(tbl.rows[i].cells[2].innerHTML)
        if ((minSelf !== 0 && selfStake < minSelf) || (maxSelf !== 0 && selfStake >= maxSelf)) {
            tbl.deleteRow(i)
            --i
            continue
        }
        const pct = parseInt(tbl.rows[i].cells[4].innerHTML)
        if (missed !== 0 && pct >= missed) {
            tbl.deleteRow(i)
            --i
            continue
        }
        const voted = parseInt(tbl.rows[i].cells[7].innerHTML)
        if (voted === 0 && gov === true) {
            tbl.deleteRow(i)
            --i
        }
    }
}

const testData = [
    {"name": "test123", "stake": 123123123, "self_stake": 123123, "missed": 11, "missed_pct": 4, "unclaimed_rew": 22222, "unclaimed_com": 33333, "votes": 22},
    {"name": "test223", "stake": 3123123, "self_stake": 3123, "missed": 1, "missed_pct": 1, "unclaimed_rew": 222, "unclaimed_com": 333, "votes": 2},
    {"name": "test323", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test33", "stake": 1123, "self_stake": 1, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test123", "stake": 123123123, "self_stake": 123123, "missed": 11, "missed_pct": 4, "unclaimed_rew": 22222, "unclaimed_com": 33333, "votes": 22},
    {"name": "test123", "stake": 123123123, "self_stake": 123123, "missed": 11, "missed_pct": 4, "unclaimed_rew": 22222, "unclaimed_com": 33333, "votes": 22},
    {"name": "test223", "stake": 3123123, "self_stake": 3123, "missed": 1, "missed_pct": 1, "unclaimed_rew": 222, "unclaimed_com": 333, "votes": 2},
    {"name": "test323", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test33", "stake": 1123, "self_stake": 1, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test223", "stake": 3123123, "self_stake": 3123, "missed": 1, "missed_pct": 1, "unclaimed_rew": 222, "unclaimed_com": 333, "votes": 2},
    {"name": "test323", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
    {"name": "test33", "stake": 1123, "self_stake": 1, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0},
]