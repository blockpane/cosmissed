const infoCircle = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-info-circle" viewBox="0 0 16 16">
  <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
  <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
</svg>`

const warningTriangle = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="darkorange" class="bi bi-exclamation-triangle" viewBox="0 0 16 16">
  <path d="M7.938 2.016A.13.13 0 0 1 8.002 2a.13.13 0 0 1 .063.016.146.146 0 0 1 .054.057l6.857 11.667c.036.06.035.124.002.183a.163.163 0 0 1-.054.06.116.116 0 0 1-.066.017H1.146a.115.115 0 0 1-.066-.017.163.163 0 0 1-.054-.06.176.176 0 0 1 .002-.183L7.884 2.073a.147.147 0 0 1 .054-.057zm1.044-.45a1.13 1.13 0 0 0-1.96 0L.165 13.233c-.457.778.091 1.767.98 1.767h13.713c.889 0 1.438-.99.98-1.767L8.982 1.566z"/>
  <path d="M7.002 12a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 5.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 5.995z"/>
</svg>`

const slashCircle = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="red" class="bi bi-slash-circle" viewBox="0 0 16 16">
  <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
  <path d="M11.354 4.646a.5.5 0 0 0-.708 0l-6 6a.5.5 0 0 0 .708.708l6-6a.5.5 0 0 0 0-.708z"/>
</svg>`

function drawValidators(vals) {
    document.getElementById("validators").innerHTML = null
    vals.forEach ((v) => {
        let redFlag = false
        let avoid = false
        let row = document.createElement("tr")
        let name = row.insertCell(-1)
        let stake = row.insertCell(-1)
        stake.innerHTML = v["stake"].toLocaleString()
        let selfStake = row.insertCell(-1)
        selfStake.innerHTML = v["self_stake"].toLocaleString()
        if (parseFloat(v["self_stake"]) / parseFloat(v["stake"]) < 0.005) {
            selfStake.className = "text-danger"
            redFlag = true
        } else if (parseFloat(v["self_stake"]) / parseFloat(v["stake"]) <= 0.01) {
            selfStake.className = "text-warning"
        } if (parseFloat(v["self_stake"]) / parseFloat(v["stake"]) >= 0.1) {
            selfStake.className = "text-success"
        }
        let missedCount = row.insertCell(-1)
        missedCount.innerHTML = v["missed"].toLocaleString()
        let missed = row.insertCell(-1)
        missed.innerHTML = v["missed_pct"].toLocaleString()
        if (parseInt(v["missed_pct"]) < 1) {
            missed.className = "text-success"
            missedCount.className = "text-success"
        } else if (parseInt(missed.innerHTML) >= 25) {
            redFlag = true
            missed.className = "text-danger"
            missedCount.className = "text-danger"
        } else if (parseInt(missed.innerHTML) >= 5) {
            missed.className = "text-warning"
            missedCount.className = "text-warning"
        }
        row.insertCell(-1).innerHTML = v["unclaimed_rew"].toLocaleString()
        row.insertCell(-1).innerHTML = v["unclaimed_com"].toLocaleString()
        let vote = row.insertCell(-1)
        vote.innerHTML = v["votes"].toLocaleString()
        if (parseInt(v["votes"]) < 1) {
            vote.className = "text-warning"
        }
        let com = row.insertCell(-1)
        com.innerHTML = v["commission"].toLocaleString()
        if (parseInt(v["commission"]) <= 1) {
            com.className = "text-warning"
        } else if (parseInt(v["commission"]) === 100) {
            com.className = "text-danger fw-bold text-decoration-underline"
            avoid = true
            redFlag = true
        } else if (parseInt(v["commission"]) > 10) {
            com.className = "text-danger fw-bold"
            redFlag = true
        } else if (parseInt(v["commission"]) <= 5) {
            com.className = "text-success"
        }
        if (avoid === true){
            name.innerHTML = `<button class="btn-light">${slashCircle}</button> &nbsp;` + v["name"]
        } else if (redFlag === true) {
            name.innerHTML = `<button class="btn-light">${warningTriangle}</button> &nbsp;` + v["name"]
        } else {
            name.innerHTML = `<button class="btn-light">${infoCircle}</button> &nbsp;` + v["name"]
        }
        name.className = "text-start"
        document.getElementById("validators").appendChild(row)
    })
}

function filterTable(minStake = 0, maxStake = 0, minSelf = 0, maxSelf = 0, missed = 0, gov = false) {
    let tbl = document.getElementById("sortTable")
    for (let i = 0; i < tbl.rows.length; i++) {
        const stake = parseInt(tbl.rows[i].cells[1].innerHTML.replaceAll(',', ''))
        if ((stake < minStake && minStake !== 0) || (stake >= maxStake && maxStake !== 0)) {
            tbl.deleteRow(i)
            --i
            continue
        }
        const selfStake = parseInt(tbl.rows[i].cells[2].innerHTML.replaceAll(',', ''))
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
        const voted = parseInt(tbl.rows[i].cells[7].innerHTML.replaceAll(',', ''))
        if (voted === 0 && gov === true) {
            tbl.deleteRow(i)
            --i
        }
    }
}

const testData = [
    {"name": "test123", "stake": 1233123, "self_stake": 323123, "missed": 11, "missed_pct": 4, "unclaimed_rew": 22222, "unclaimed_com": 33333, "votes": 22, "commission": 1},
    {"name": "test223", "stake": 3123123, "self_stake": 3123, "missed": 1, "missed_pct": 1, "unclaimed_rew": 1222, "unclaimed_com": 333, "votes": 2, "commission": 2},
    {"name": "test323", "stake": 12323, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 5},
    {"name": "test", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 0.5},
    {"name": "test33", "stake": 1123, "self_stake": 100, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 7},
    {"name": "test123", "stake": 1233123, "self_stake": 123123, "missed": 11, "missed_pct": 4, "unclaimed_rew": 22222, "unclaimed_com": 33333, "votes": 22, "commission": 5},
    {"name": "test123", "stake": 12312313, "self_stake": 123123, "missed": 11, "missed_pct": 4, "unclaimed_rew": 22222, "unclaimed_com": 33333, "votes": 22, "commission": 5},
    {"name": "test223", "stake": 3123, "self_stake": 3123, "missed": 111, "missed_pct": 11.2, "unclaimed_rew": 222, "unclaimed_com": 333, "votes": 2, "commission": 5},
    {"name": "test323", "stake": 123123, "self_stake": 1259, "missed": 3005, "missed_pct": 32.5, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 3},
    {"name": "test", "stake": 123123, "self_stake": 12304, "missed": 245, "missed_pct": 23.5, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 100},
    {"name": "test33", "stake": 1123, "self_stake": 1000, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 10},
    {"name": "test223", "stake": 3123123, "self_stake": 3123, "missed": 1, "missed_pct": 0.14, "unclaimed_rew": 222, "unclaimed_com": 333, "votes": 2, "commission": 5},
    {"name": "test323", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 14},
    {"name": "test", "stake": 123123, "self_stake": 123, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 3.5},
    {"name": "test", "stake": 123123, "self_stake": 112304, "missed": 2145, "missed_pct": 23.5, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 100},
    {"name": "test3", "stake": 111123, "self_stake": 10010, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 10},
    {"name": "testz223", "stake": 3123123, "self_stake": 3123, "missed": 1, "missed_pct": 0.14, "unclaimed_rew": 222, "unclaimed_com": 333, "votes": 2, "commission": 5},
    {"name": "test3d23", "stake": 123123, "self_stake": 1023, "missed": 0, "missed_pct": 0, "unclaimed_rew": 0, "unclaimed_com": 0, "votes": 0, "commission": 10},
]