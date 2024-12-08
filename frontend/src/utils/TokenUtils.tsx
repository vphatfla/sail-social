// eslint-disable-next-line @typescript-eslint/no-explicit-any
const decodeToken = (token: string): any | null => {
    try {
        const rawPL = token.split('.')[1];
        const decodeddPL = atob(rawPL);
        return JSON.parse(decodeddPL);
    } catch (error) {
        console.log("Failed to decode token ", error);
        return null;
    }
}

const saveDataLocalStorage = (decodedPL: JSON) => {
    if (decodedPL !== null) {
        Object.entries(decodedPL).forEach(([key, value]) => {
            localStorage.setItem(key, value + '')
        })
    }
}
export { decodeToken, saveDataLocalStorage };