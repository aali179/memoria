let tabButtons = [document.querySelectorAll('.drop-down')]

// Toggle content when tabs are clicked
function showCanvasContent(id) {  
    for (let i=0; i<tabButtons[0].length; i++){
        console.log(tabButtons[0][i].id)
        if (tabButtons[0][i].id == id) {
            document.getElementById(id.replace('btn', 'content')).style.display = "block";
        } else {
            document.getElementById(tabButtons[0][i].id.replace('btn', 'content')).style.display = "none";
        }
    }
 }

let imagesEl = document.getElementById("images-btn")
imagesEl.addEventListener("click", function() {
    showCanvasContent("images-btn")
})

let textEl = document.getElementById("text-btn")
textEl.addEventListener("click", function() {
    showCanvasContent("text-btn")
})

let musicEl = document.getElementById("music-btn")
musicEl.addEventListener("click", function() {
    showCanvasContent("music-btn")
})

let mapsEl = document.getElementById("maps-btn")
mapsEl.addEventListener("click", function() {
    showCanvasContent("maps-btn")
})
