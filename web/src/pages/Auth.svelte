<script>
  import { onMount } from "svelte";

  let authURL = "";

  let code = "";

  onMount(async () => {
    const res = await fetch(`http://localhost:8000/initialauth`);
    authURL = await res.json();
  });

  async function postData(url = "", data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
      method: "POST", // *GET, POST, PUT, DELETE, etc.
      mode: "cors", // no-cors, *cors, same-origin
      cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
      credentials: "same-origin", // include, *same-origin, omit
      headers: {
        "Content-Type": "application/json"
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      redirect: "follow", // manual, *follow, error
      referrerPolicy: "no-referrer", // no-referrer, *client
      body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return await response.json(); // parses JSON response into native JavaScript objects
  }

  function handleClick() {
    postData("http://localhost:8000/completeauth", { code: code }).then(
      data => {
        console.log(data); // JSON data parsed by `response.json()` call
      }
    );
  }
</script>

<style>
  .main-container {
    padding: 30px;
    max-width: 700px;
  }

  .auth-button {
    padding: 20px;
    background: lightcoral;
    color: white;
    font-family: "Courier New", Courier, monospace;
  }

  .complete-button {
    padding: 20px;
    margin-top: 30px;
    background: lightcoral;
    color: white;
    font-family: "Courier New", Courier, monospace;
  }

  .complete-button:hover {
    cursor: pointer;
  }

  .code-input {
    padding: 20px;
  }
</style>

<div class="main-container">
  <button class="auth-button">HI .</button>

  <p>
    Visit
    <a href={authURL} target="_blank">Google Auth</a>
    complete the authentication and then paste the auth code given below
  </p>

  <div>
    <input bind:value={code} class="code-input" />
  </div>

  <button on:click={handleClick} class="complete-button">Complete Auth</button>
</div>
