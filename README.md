# [WailsV2] G-Earth Extension Starter Kit

Welcome to the **[WailsV2] G-Earth Extension Starter Kit!** This demo project showcases how to use **GoEarth**, **Wails v2**, **Vue**, and **TailwindCSS** to build custom extensions with a full graphical user interface (GUI).

## About

This project serves as a template for developing extensions that interact with G-Earth, providing real-time updates and a seamless user experience. It is designed to help you get started with extension development and learn how to build GUIs that communicate effectively with backend services.

### Key Features

- **Cross-Platform GUI**: Utilize GoEarth with Wails to create applications that run on multiple platforms.
- **Data Communication**: Pass data between the Go backend and the Vue frontend effortlessly.
- **Real-Time Interaction**: Engage with G-Earth through real-time events and updates.

## Live Development

To run the project in live development mode, execute the following command in your project directory:

```bash
wails dev
```

This will start a Vite development server, enabling fast hot reloading of frontend changes. For browser development, you can connect to the dev server at [http://localhost:34115](http://localhost:34115) to access your Go methods from the browser's developer tools.

## Building for Production

To create a production-ready package, run:

```bash
wails build
```

This command will generate a redistributable version of your application.

## Getting Started

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/JTD420/G-Starter-Kit-V2.git
   cd G-Starter-Kit-V2
   ```

2. Install the necessary dependencies:
   ```bash
   npm install
   ```

3. Follow the Live Development instructions above to start building your extension.

## Example Usage

This demo application serves as a **basic chat logger**, which tracks users entering and leaving the room, as well as messages sent by users. It logs the username of the sender along with the message content in real time. 

To ensure you capture all user interactions when starting the extension while already in a room, make sure to reload the room to fetch the list of all current users.

### User Interaction Handling

The main application file demonstrates how to manage user interactions and display messages received from the backend. Reactive variables `username` and `message` are utilized to show incoming messages and user events dynamically.

### Sample Data Display

Here’s an example of how the application displays data passed from the backend:

```json
{
  "User": "{{ username }}",
  "Message": "{{ message }}"
}
```

- The `User` field reflects the name of the user interacting in the chat (either sending a message or entering/leaving the room).
- The `Message` field provides the content of the user's message or indicates if a user has entered or left the room.

### Data Flow Explanation

1. **User Entry**: When a user enters the room, the `handleUsers` function logs the entry event and emits an event with the username and a message indicating they have entered.
2. **Chat Messages**: When a message is sent, the `handleChat` function captures the username and the message content, emitting a `newMessage` event to the frontend.
3. **User Exit**: If a user leaves the room, the `handleRemoveUser` function captures this event and emits an event to indicate the user has left.

This structure allows for real-time updates in the chat logger, providing a smooth user experience as users enter, leave, and communicate in the room.


## Want to Learn More?

- [Check out GoEarth](https://github.com/xabbo/goearth) to explore its capabilities.
- [Visit Wails Documentation](https://wails.io/docs/introduction) for more information about building applications with Wails.

---

Thank you for using the [WailsV2] G-Earth Extension Starter Kit! I hope this template helps you create some amazing extensions! If you have any questions or feedback, feel free to open an issue in this repository.
