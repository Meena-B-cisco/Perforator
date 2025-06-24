export const openLink = (url?: string) => {
    if (url) {
        window
            .open(url, '_blank')
            ?.focus();
    }
};
