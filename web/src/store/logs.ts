import { writable, derived, readable } from "svelte/store";
let logs = [
];

const { subscribe, set, update } = writable(logs);

const addLog = (log) =>
  update((logs) => {
    return [...logs, ...log];
  });

export default {
  subscribe,
  addLog,
};
