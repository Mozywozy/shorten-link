async function shortenURL() {
    const longURL = document.getElementById('longURL').value;
    const resultDiv = document.getElementById('result');
    
    if (!longURL) {
        resultDiv.innerHTML = '<p>Please enter a URL.</p>';
        return;
    }

    try {
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ long_url: longURL })
        });
        const data = await response.json();

        if (response.ok) {
            resultDiv.innerHTML = `<p>Shortened URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a></p>`;
        } else {
            resultDiv.innerHTML = `<p>Error: ${data.error}</p>`;
        }
    } catch (error) {
        resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
    }
}
