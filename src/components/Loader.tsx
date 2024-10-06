import { CircularProgress } from "@mui/material";

const Loader = () => {
    return (
        <main className="flex items-center justify-center h-screen">
            <CircularProgress color="error" />
        </main>
    )
}

export default Loader;
