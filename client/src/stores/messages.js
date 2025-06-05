export const useMessagesStore = defineStore("messages", () => {
    const queue = ref([]);
    function add(message) {
        queue.push(message);
    }

    return { queue, add };
});
