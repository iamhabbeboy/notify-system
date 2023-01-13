import { onMount } from "svelte";
import { writable } from "svelte/store";
let logs = [
  // {
  //   id: 1,
  //   serviceName: "slack",
  //   environment: "production",
  //   message: "error occured at line 1092",
  //   date: "2023-03-02",
  // },
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
