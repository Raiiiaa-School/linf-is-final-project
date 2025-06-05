import { onMounted, onUnmounted, ref } from "vue";

export function useQueryParams() {
    const queryParams = ref({});

    const parseQueryParams = () => {
        const params = new URLSearchParams(window.location.search);
        const newParams = {};
        for (const [key, value] of params.entries()) {
            newParams[key] = value;
        }
        queryParams.value = newParams;
    };

    const updateURL = (newParams, replace = false) => {
        const currentParams = new URLSearchParams(window.location.search);
        for (const key in newParams) {
            if (newParams[key] === null || newParams[key] === undefined) {
                currentParams.delete(key);
            } else {
                currentParams.set(key, newParams[key]);
            }
        }

        const queryString = currentParams.toString();
        const newURL = `${window.location.pathname}${
            queryString ? "?" + queryString : ""
        }${window.location.hash}`;

        if (replace) {
            window.history.replaceState({}, "", newURL);
        } else {
            window.history.pushState({}, "", newURL);
        }

        parseQueryParams();
    };

    const handlePopState = () => {
        parseQueryParams();
    };

    onMounted(() => {
        parseQueryParams();
        window.addEventListener("popstate", handlePopState);
    });

    onUnmounted(() => {
        window.removeEventListener("popstate", handlePopState);
    });

    const setQueryParams = (params, replace = false) => {
        const currentParams = { ...queryParams.value };
        const updatedParams = { ...currentParams, ...params };
        updateURL(updatedParams, replace);
    };

    const getQueryParam = (key) => {
        return queryParams.value[key] || null;
    };

    const removeQueryParam = (keys, replace = false) => {
        const keysToRemove = Array.isArray(keys) ? keys : [keys];
        const newParams = { ...queryParams.value };
        keysToRemove.forEach((k) => {
            delete newParams[k];
        });
        updateURL(newParams, replace);
    };

    return {
        queryParams,
        setQueryParams,
        getQueryParam,
        removeQueryParam,
    };
}
