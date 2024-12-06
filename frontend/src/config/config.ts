interface Config {
    SERVER_URL: string
}

const getConfig = (): Config => {
    const { VITE_IS_LOCAL, VITE_LOCAL_SERVER_URL, VITE_PROD_SERVER_URL } = import.meta.env;

    if (!VITE_IS_LOCAL) throw new Error("IS_LOCAL not defined");
    if (!VITE_LOCAL_SERVER_URL) throw new Error("LOCAL_SERVER_URL not defined")
    if (!VITE_PROD_SERVER_URL) throw new Error("PROD_SERVER_URL not defined");

    return {
        SERVER_URL: (VITE_IS_LOCAL === "TRUE") ? VITE_LOCAL_SERVER_URL : VITE_PROD_SERVER_URL,
    };
}

const config = getConfig();

export default config;