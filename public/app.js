new Vue({ // Mount vue object to a div with the id of #app
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function() { // Vue function initializes when the app launches
        var self = this;
        this.ws = new WebSocket('wss://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) { // Takes a function to handle incoming messages as JSON string and parse them as object literals
            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                    + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>'; // Parse emojis using EmojiOne library

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },


    methods: { // Defines the functions we want our app to have access to in the scope of the Vue app
        send: function () { // Function for sending messages to server
            if (this.newMsg != '') { // Make sure message isn't blank
                this.ws.send( 
                    JSON.stringify({ // Format message as object and stringify it
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text(), // Strip out html to prevent injection
                        time: this.time,
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () { // Users joining server must enter credentials 
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text(); // Strip out html 
            this.username = $('<p>').html(this.username).text();
            this.joined = true; // Set joined to true if above conditions are met
        },

        gravatarURL: function(email) { // Grabs avatar using user's email address
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email); // One-way encryption to protect user email address submissions
        }
    }
});