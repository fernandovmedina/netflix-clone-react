import ProfileCards from "./layouts/ProfileCards.tsx";

const Browse = () => {
    return (
        <main>
            <section
                className="bg-gradient-to-b from-black to-gray-950 text-white h-[100vh] flex items-center justify-center"
            >
                <div className="text-center">
                    <h1 className="font-bold text-5xl">Who's watching now?</h1>
                    <ProfileCards />
                    <a
                        href="/ManageProfiles"
                        className="border-2 border-gray-400 hover:border-white hover:text-white px-3 py-2 text-gray-400"
                    >Manage profiles
                    </a>
                </div>
            </section>
        </main>
    )
}

export default Browse;
