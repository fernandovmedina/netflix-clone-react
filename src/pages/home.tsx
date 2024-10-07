import FooterHome from "../components/FooterHome";
import NavHome from "../components/NavHome";
import { movies, series } from "./series";

const Home = () => {
    const renderItems = (items: Record<string, string>, title: string) => (
        <div className="pb-10">
            <h1 className="text-xl mb-3">{title}</h1>
            <div className="flex">
                {Object.entries(items).map(([key, value]) => (
                    <img className="px-1 w-72" key={key} src={value} />
                ))}
            </div>
        </div>
    );

    return (
        <main>
            <NavHome />

            <section className="bg-black text-white px-16">
                {renderItems(movies, "Movies")}
                {renderItems(series, "Series")}
            </section>

            <FooterHome />
        </main>
    );
};

export default Home;
