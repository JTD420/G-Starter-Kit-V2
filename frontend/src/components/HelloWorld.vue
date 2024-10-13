<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import logo from '/src/assets/images/logo.svg';
import beams from '/src/assets/images/beams.jpg';

const username = ref(""); // Create a reactive variable for username
const message = ref("");   // Create a reactive variable for message

// Function to handle incoming messages
function handleNewMessage(data) {
  username.value = data.username || "Guest"; // Fallback to "Guest" if username is not present
  message.value = data.message || "No new messages."; // Fallback message
}

// Function to handle incoming user event
function handleUserEvent(data) {
  username.value = data.username || "Guest"; // Fallback to "Guest" if username is not present
  message.value = data.message
}

// Listen for the newMessage event
onMounted(() => {
  window.runtime.EventsOn("newMessage", handleNewMessage);
  window.runtime.EventsOn("UserEvent", handleUserEvent);
});

// Clean up event listener on unmount
onUnmounted(() => {
  window.runtime.EventsRemove("newMessage", handleNewMessage);
  window.runtime.EventsRemove("UserEvent", handleUserEvent);
});
</script>

<template>
  <main class="relative flex min-h-screen flex-col justify-center overflow-hidden bg-gray-50 py-6 sm:py-12">
    <img :src="beams" alt="" class="absolute top-1/2 left-1/2 max-w-none -translate-x-1/2 -translate-y-1/2" width="1308" />
    <div class="absolute inset-0 bg-[url(./assets/images/grid.svg)] bg-center [mask-image:linear-gradient(180deg,white,rgba(255,255,255,0))]"></div>
    <div class="relative bg-white px-6 pt-10 pb-8 shadow-xl ring-1 ring-gray-900/5 sm:mx-auto sm:max-w-lg sm:rounded-lg sm:px-10">
      <div class="mx-auto max-w-md">
        <img :src="logo" class="h-6" alt="GoEarth Starter Kit" />
        <div class="divide-y divide-gray-300/50">
          <div class="space-y-6 py-8 text-base leading-7 text-gray-600">
            <p>Welcome to the G-Earth Extensions Starter Kit! This demo showcases how to use GoEarth, Wails V2, Vue, and TailwindCSS to build custom extensions with a full GUI.</p>
            <ul class="space-y-4">
              <li class="flex items-center">
                <svg class="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="11" />
                  <path d="m8 13 2.165 2.165a1 1 0 0 0 1.521-.126L16 9" fill="none" />
                </svg>
                <p class="ml-4 mt-8">Using GoEarth with Wails to build a cross-platform GUI</p>
              </li>
              <li class="flex items-center">
                <svg class="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="11" />
                  <path d="m8 13 2.165 2.165a1 1 0 0 0 1.521-.126L16 9" fill="none" />
                </svg>
                <p class="ml-4">Passing data from the Go backend to the Vue frontend</p>
              </li>
              <li class="flex items-center">
                <svg class="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="11" />
                  <path d="m8 13 2.165 2.165a1 1 0 0 0 1.521-.126L16 9" fill="none" />
                </svg>
                <p class="ml-4">Real-time interaction with G-Earth</p>
              </li>
            </ul>
            <p>This template is perfect for getting started with extension development and building a GUI that interacts with G-Earth!</p>
          </div>

          <!-- Game Data Display -->
          <div class="pt-8">
            <h3 class="text-lg font-semibold text-gray-900">Game Data</h3>
            <div class="mt-4 p-4 bg-gray-100 rounded-lg shadow-inner">
              <p class="text-sm text-gray-600">Here is an example of how to display data passed from the backend:</p>
              <pre class="mt-2 p-2 bg-gray-800 text-white rounded">
{
  "User": "{{ username }}",
  "Message": "{{ message }}"
}
              </pre>
            </div>
          </div>

          <div class="pt-8 text-base font-semibold leading-7">
            <p class="text-gray-900">Want to learn more?</p>
            <p>
              <a href="https://github.com/xabbo/goearth" class="text-sky-500 hover:text-sky-600">Check out GoEarth &rarr;</a>
            </p>
            <p>
              <a href="https://wails.io/docs/introduction" class="text-sky-500 hover:text-sky-600">Check out Wails &rarr;</a>
            </p>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
