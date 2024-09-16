import { useEffect, useRef } from "react";
import FooterSign from "../../components/FooterSign";
import NavbarOut from "../../components/NavbarOut";

const OxxoCodeGenerator = () => {
    const backRef = useRef<HTMLHeadingElement>(null);
    const email: string | null = new URLSearchParams(window.location.search).get("email");
    const password: string | null = new URLSearchParams(window.location.search).get("password");
    const plan: string | null = new URLSearchParams(window.location.search).get("plan");

    useEffect(() => {
        const backElement = backRef.current;
        if (backElement) {
            const handleClick = () => {
                window.location.href = `/signup/cashPaymentOption?email=${email}&password=${password}&plan=${plan}`;
            };
            backElement.addEventListener('click', handleClick);

            return () => {
                backElement.removeEventListener('click', handleClick);
            };
        }
    }, [email, password, plan]);

    return (
        <main>
            <NavbarOut />

            <section className="text-center my-20">
                <h2 ref={backRef} className="text-blue-500 hover:underline hover:underline-offset-1">Go Back</h2>
                <h1 className="my-2 font-bold text-3xl">Random Code Generator</h1>
                <p className="text-xl">The generated codes will be stored in our database and you are going to be able to use it as a payment method</p>
                <img src="" alt="bar_code" />
                <h4 id="code"></h4>
                <button className="bg-red-600 px-5 py-3 rounded mt-10 text-white hover:bg-red-900 font-bold">Generate</button>
            </section>

            <FooterSign />
        </main>
    );
}

export default OxxoCodeGenerator;
