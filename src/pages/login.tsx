import { useEffect } from 'react';

const Login = () => {
    useEffect(() => {
        const rmb = document.getElementById("remember");
        const btn = document.getElementById("btn");
        const inf = document.getElementById("inf");
        const hidden = document.getElementById("hidden");

        if (rmb && btn) {
            rmb.addEventListener("click", () => {
                btn.checked = !btn.checked;
            });
        }

        if (inf && hidden) {
            inf.addEventListener("click", () => {
                hidden.style.display = "block";
                inf.style.display = "none";
            });
        }

        // Cleanup event listeners when the component is unmounted
        return () => {
            if (rmb && btn) {
                rmb.removeEventListener("click", () => {
                    btn.checked = !btn.checked;
                });
            }
            if (inf && hidden) {
                inf.removeEventListener("click", () => {
                    hidden.style.display = "block";
                    inf.style.display = "none";
                });
            }
        };
    }, []); // El array vacío [] asegura que este efecto solo se ejecute una vez, después del primer renderizado

    return (
        <main>
            <div className="bg-[url(/MX-es-20240603-popsignuptwoweeks-perspective_alpha_website_medium.jpg)] bg-auto">
                <nav className="w-full flex py-5 px-40">
                    <div className="w-1/2">
                        <a href="/">
                            <svg
                                className="w-40 default-ltr-cache-1d568uk ev1dnif2"
                                viewBox="0 0 111 30"
                                version="1.1"
                                xmlns="http://www.w3.org/2000/svg"
                                xmlns:xlink="http://www.w3.org/1999/xlink"
                                aria-hidden="true"
                                role="img"
                            >
                                <g>
                                    <path
                                        fill="red"
                                        d="M105.06233,14.2806261 L110.999156,30 C109.249227,29.7497422 107.500234,29.4366857 105.718437,29.1554972 L102.374168,20.4686475 L98.9371075,28.4375293 C97.2499766,28.1563408 95.5928391,28.061674 93.9057081,27.8432843 L99.9372012,14.0931671 L94.4680851,-5.68434189e-14 L99.5313525,-5.68434189e-14 L102.593495,7.87421502 L105.874965,-5.68434189e-14 L110.999156,-5.68434189e-14 L105.06233,14.2806261 Z M90.4686475,-5.68434189e-14 L85.8749649,-5.68434189e-14 L85.8749649,27.2499766 C87.3746368,27.3437061 88.9371075,27.4055675 90.4686475,27.5930265 L90.4686475,-5.68434189e-14 Z M81.9055207,26.93692 C77.7186241,26.6557316 73.5307901,26.4064111 69.250164,26.3117443 L69.250164,-5.68434189e-14 L73.9366389,-5.68434189e-14 L73.9366389,21.8745899 C76.6248008,21.9373887 79.3120255,22.1557784 81.9055207,22.2804387 L81.9055207,26.93692 Z M64.2496954,10.6561065 L64.2496954,15.3435186 L57.8442216,15.3435186 L57.8442216,25.9996251 L53.2186709,25.9996251 L53.2186709,-5.68434189e-14 L66.3436123,-5.68434189e-14 L66.3436123,4.68741213 L57.8442216,4.68741213 L57.8442216,10.6561065 L64.2496954,10.6561065 Z M45.3435186,4.68741213 L45.3435186,26.2498828 C43.7810479,26.2498828 42.1876465,26.2498828 40.6561065,26.3117443 L40.6561065,4.68741213 L35.8121661,4.68741213 L35.8121661,-5.68434189e-14 L50.2183897,-5.68434189e-14 L50.2183897,4.68741213 L45.3435186,4.68741213 Z M30.749836,15.5928391 C28.687787,15.5928391 26.2498828,15.5928391 24.4999531,15.6875059 L24.4999531,22.6562939 C27.2499766,22.4678976 30,22.2495079 32.7809542,22.1557784 L32.7809542,26.6557316 L19.812541,27.6876933 L19.812541,-5.68434189e-14 L32.7809542,-5.68434189e-14 L32.7809542,4.68741213 L24.4999531,4.68741213 L24.4999531,10.9991564 C26.3126816,10.9991564 29.0936358,10.9054269 30.749836,10.9054269 L30.749836,15.5928391 Z M4.78114163,12.9684132 L4.78114163,29.3429562 C3.09401069,29.5313525 1.59340144,29.7497422 0,30 L0,-5.68434189e-14 L4.4690224,-5.68434189e-14 L10.562377,17.0315868 L10.562377,-5.68434189e-14 L15.2497891,-5.68434189e-14 L15.2497891,28.061674 C13.5935889,28.3437998 11.906458,28.4375293 10.1246602,28.6868498 L4.78114163,12.9684132 Z"
                                    ></path>
                                </g>
                            </svg>
                        </a>
                    </div>
                    <div className="w-1/2"></div>
                </nav>

                <section className="flex items-center justify-center pb-10">
                    <div className="auth_bg py-10 text-white px-14 rounded w-[35%]">
                        <form action="" method="post" className="flex flex-col">
                            <h1 className="font-extrabold text-3xl pb-4">Login</h1>
                            <input
                                type="email"
                                placeholder="Email"
                                className="mb-3 bg-transparent border-2 border-gray-500 rounded-md px-3 py-2"
                            />
                            <input
                                type="password"
                                placeholder="Password"
                                className="mb-3 bg-transparent border-2 border-gray-500 rounded-md px-3 py-2"
                            />
                            <button
                                type="submit"
                                className="bg-red-700 hover:bg-red-800 py-2 rounded"
                            >Login</button>
                            <h1 className="my-3 mx-auto">O</h1>
                            <a
                                href=""
                                className="text-center bg-gray-500 hover:bg-gray-600 w-full py-2 rounded mb-3"
                            >Use a login code</a>
                            <a
                                href=""
                                className="mx-auto hover:underline hover:underline-offset-1"
                            >Forgot your password?</a>
                            <div className="flex py-2">
                                <input type="checkbox" checked id="btn" />
                                <button type="button" id="remember" className="ml-3"
                                >Remember me</button>
                            </div>
                            <div className="flex py-2">
                                <p>First time on Netflix?</p>
                                <a
                                    href="/"
                                    className="font-extrabold ml-2 hover:underline hover:underline-offset-1"
                                >Subscribe now</a>
                            </div>
                            <p className="text-sm text-gray-500">
                                This page is protected for Google reCAPTCHA to
                                verify you are not a robot. <button
                                    type="button"
                                    className="text-blue-600 hover:underline hover:underline-offset-1"
                                    id="inf">More info</button>
                            </p>
                            <div className="hidden mt-2" id="hidden">
                                <p className="text-sm text-gray-500">
                                    The information collected by Google reCAPTCHA is
                                    subject to Google's <a
                                        className="text-blue-600 hover:underline hover:underline-offset-1"
                                        href="https://policies.google.com/privacy"
                                    >Privacy Policy</a> and <a
                                        href="https://policies.google.com/terms"
                                        className="text-blue-600 hover:underline hover:underline-offset-1"
                                    >Terms of Service</a>, and is used to provide, maintain and improve
                                    the reCAPTCHA service, as well as for general
                                    security purposes (Google does not use it to
                                    personalize advertising).
                                </p>
                            </div>
                        </form>
                    </div>
                </section>
            </div>

            <footer className="w-full bg-black text-gray-400 px-40 py-12">
                <h4 className="font-normal mb-4">
                    Questions? Call to <a href="tel:800 953 1430">800 953 1430</a>
                </h4>
                <div className="w-full flex text-sm underline underline-offset-1 pb-4">
                    <div className="w-1/4 flex flex-col">
                        <a href="" className="hover:text-gray-200">Frecuent questions</a>
                        <a href="" className="hover:text-gray-200">Privacy</a>
                    </div>
                    <div className="w-1/4 flex flex-col">
                        <a href="" className="hover:text-gray-200">Help center</a>
                        <a href="" className="hover:text-gray-200"
                        >Cookies preferences</a>
                    </div>
                    <div className="w-1/4 flex flex-col">
                        <a href="" className="hover:text-gray-200">Netflix store</a>
                        <a href="" className="hover:text-gray-200"
                        >Corporative information</a>
                    </div>
                    <div className="w-1/4 flex flex-col">
                        <a href="" className="hover:text-gray-200">Terms of use</a>
                    </div>
                </div>
                <button
                    className="flex items-center mx-2 bg-black mb-4 text-white border border-gray-500 px-2 py-1 rounded"
                >
                    <img
                        src="/icons8-translate-google-64.png"
                        alt="google_translate_img"
                        className="w-6"
                    />
                    <select
                        name="languages"
                        id="languages"
                        className="bg-black text-white"
                    >
                        <option value="english">English</option>
                        <option value="spanish">Spanish</option>
                    </select>
                </button>
                <p>Netflix</p>
            </footer>
        </main>
    );
}

export default Login;
