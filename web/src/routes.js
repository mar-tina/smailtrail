import Home from "./pages/Home.svelte";
import Auth from "./pages/Auth.svelte";
import NotFound from "./pages/NotFound.svelte";
import Messages from "./pages/Messages.svelte";
import Thread from './pages/Thread.svelte';

export const routes = {
  // Exact path
  "/": Home,

  "/auth": Auth,

  "/un-sub": Home,

  "/msgs": Messages,

  "/thread/:id" : Thread,
  // Catch-all
  // This is optional, but if present it must be the last
  "*": NotFound
};
