import { writable, derived, readable } from "svelte/store";
let logs = [
  // {
  //   id: 1,
  //   serviceName: "slack",
  //   environment: "production",
  //   message: "error occured at line 1092",
  //   date: "2023-03-02",
  // },
];

const filter = writable({ service: 'all', environment: 'production', message: '', date: ''})

const { subscribe, set, update } = writable(logs);

const addLog = (log) =>
  update((logs) => {
    return [...logs, ...log];
  });

	const filtered = derived([filter, readable(logs)], ([$filter, $logs]) => {
		if ($filter.service == 'all') return $logs;
		return $logs.filter(data => (data.service === $filter.service) || ($filter.message.includes(data.message)))
	})

export default {
  subscribe,
  addLog,
  filter,
  filtered,
};
