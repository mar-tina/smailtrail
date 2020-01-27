


const Auth = {
  data() {
    return {
      loading: false,
      authURL: "",
      error: null,
      authCode: "",
      styleObject: {
        padding: "30px",
        display: "grid",
        justifyContent: "center",
        justifyItems: "center",
        alignItems: "center"
      },

      buttonObject: {
        padding: "10px",
        border: "none",
        background: "lightcoral",
        color: "white",
        fontSize: "15px",
        "&:hover": {
          cursor: "pointer"
        }
      },
      inputObject: {
        height: "40px",
        width: "150px"
      },

      signUpButton: {
        marginTop: "20px",
        padding: "10px",
        border: "none",
        background: "lightblue",
        color: "white",
        fontSize: "15px",
        "&:hover": {
          cursor: "pointer"
        }
      }
    };
  },
  template: `<div v-bind:style='styleObject'>
        <button v-bind:style='buttonObject' v-on:click='authme'> Click here to get started</button>

        <p> On Auth Complete. Copy the token result from the above URL and paste it here: </p>
        <input v-bind:style='inputObject' v-model="authCode">
        <br />
        <button v-bind:style='signUpButton'> Complete Authentication </button>
    </div>`,
  created: function() {
    this.fetchData();
  },

  methods: {
    fetchData() {
      this.error = this.post = null;
      this.loading = true;
      // replace `getPost` with your data fetching util / API wrapper
      fetch("/initialauth")
        .then(response => {
          this.loading = false;
          console.log(response);
          return response.json();
        })
        .then(myJson => {
          this.authURL = myJson;
          console.log(myJson);
        });
    },

    authme() {
      window.open(this.authURL);
    }
  }
};

module.exports = {
  Auth
};
