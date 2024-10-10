function calculate()
{
    console.log("salam")
    let number = document.getElementById("Number")
    let selects = document.getElementsByTagName("select")

    sendQuery(number.value, selects[0].value, selects[1].value)
}

async function sendQuery(num, curBase, newBase) 
{   

    let res = document.getElementById("result-text")

    const baseUrl = 'https://base-convertor.liara.run/convert-base'
    const params = new URLSearchParams({
        num: num,
        curBase: curBase,
        newBase: newBase,
    })

   const url = `${baseUrl}?${params}`;

    try
    {
        const response = await fetch(url)
        const result = await response.text();
        res.value = result
    } 
    
    catch (error)
    {
        console.error('Error:', error);
    }
}