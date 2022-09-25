async function getLinks() {
    const response = await fetch('http://localhost:5000/v1/links', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify({
            text: "melitopol",
            languages: [
                "UKR",
                "POL",
                "RUS",
                "AR"
            ]
        })
    })
    responseJson = await response.json()
    console.log(responseJson)
    for (const link of responseJson.Links) {
        console.log(link)
        // This isn't working for some reason and I don't know why
        let _ = await chrome.tabs.create({ url: link });
    }
}

chrome.action.onClicked.addListener((tab) => {
    chrome.scripting.executeScript({
    target: { tabId: tab.id },
    function: getLinks
    });
});