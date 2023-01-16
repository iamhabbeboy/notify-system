import { writable } from "svelte/store";

export const logFilter = writable({
    search: '',
    date: '',
});

export function updateSearchValue(searchValue: string): void {
    logFilter.update(filter => {
        filter.search = searchValue
        return filter
    })
}

export function updateDateValue(selectedDate: string): void {
    logFilter.update(filter => {
        filter.date = selectedDate;
        return filter;
    })
}