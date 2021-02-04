// import { format } from "url";

new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        avatar: null, // avatar address used for grabbing an avatar
        username: null, // Our username
        joined: false, // True if avatar and username have been filled in
        leaving: false, // false unless the user is leaving
        uuid: null
    },
    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            // if the message is encrypted, decrypt it before showing the message
            if (msg.encrypted == true) {
                msg.message = CryptoJS.AES.decrypt(msg.message, '394812730425442A472D2F423F452848').toString(CryptoJS.enc.Utf8)
            }
            // content for the chat messages
            self.chatContent += '<div class="chip">'
                    + '<img src="' + self.avatarURL(msg.avatar) + '">' // Avatar
                    + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
        // Event listener for when someone leaves.
        window.addEventListener("beforeunload", e => {
            // prevents the browsers default action for closing
            e.preventDefault
            e.returnValue = ""
            
            window.onunload = event => {
                // sends the information for the user leaving if they do.            
                this.ws.send(
                    JSON.stringify({
                        avatar: this.avatar,
                        username: this.username,
                        message: "USERLEAVING",
                        encrypted: false,
                        leaving: true,
                        uuid: this.uuid
                    })
                )
            }
        });
    },
    methods: {
        send: function () {
            // makes sure the message isnt blank, if not, sends the message to the server.
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        avatar: this.avatar,
                        username: this.username,
                        message: CryptoJS.AES.encrypt($('<p>').html(this.newMsg).text(),'394812730425442A472D2F423F452848', {
                                                                mode: CryptoJS.mode.CBC}).toString(), // Strip out html
                        encrypted: true,
                        leaving: false,
                        uuid: this.uuid
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },
        join: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            if (!this.uuid) {
                Materialize.toast('You must enter a room ID, If you do not have one then visit the create page!', 2000);
                return
            } 
            axios.post('/room/exists', {
                uuid: this.uuid
            }).then(exists => {
                console.log(exists)
                if (!exists.data) {
                    Materialize.toast('This room does not exist! please enter a valid room ID', 2000);
                    return
                } 
                this.ws.send(
                    JSON.stringify({
                            avatar: this.avatar,
                            username: this.username,
                            message: "",
                            encrypted: false,
                            leaving: false,
                            uuid: this.uuid
                    })
                );

                this.avatar = $('<p>').html(this.avatar).text();
                this.username = $('<p>').html(this.username).text();
                this.joined = true;
            })
        },
        avatarURL: function(avatar) {
             if (avatar.match(/\.(jpeg|jpg|gif|png)$/) != null) {
                return avatar;
            } else {
                return "avatar.jpg"
            }
        }
    }
});
