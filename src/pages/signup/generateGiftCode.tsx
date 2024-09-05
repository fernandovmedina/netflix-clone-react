import { useEffect } from 'react';
import FooterSign from '../../components/FooterSign.tsx'
import NavbarSign from '../../components/NavbarSign.tsx'

const GenerateGiftCode = () => {
    useEffect(() => {
        const email = new URLSearchParams(window.location.search).get("email");
        const password = new URLSearchParams(window.location.search).get("password");
        const plan = new URLSearchParams(window.location.search).get("plan");
        const back: any = document.getElementById('back');

        if (back) {
            back.addEventListener('click', () => {
                window.location.href = `/signup/giftoption?email=${email}&password=${password}&plan=${plan}`
            })
        }

        return () => {
            if (back) {
                back.removeEventListener('click', () => {
                    window.location.href = `/signup/giftoption?email=${email}&password=${password}&plan=${plan}`
                })
            }
        }
    }, [])

    const request = async () => {
        try {
            const response = await fetch("http://127.0.0.1:8040/microservice/gift_code", {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            const data = await response.json();
            const giftCodeElement = document.getElementById("gift_code");

            if (giftCodeElement) {
                giftCodeElement.textContent = data.body;
            }
        } catch (error) {
            console.error('Error while procesing the request:', error);
        }
    }

    return (
        <main>
            <div>
                <NavbarSign />
            </div>

            <section className='w-full flex items-center justify-center my-20'>
                <div className='flex flex-col items-center'>
                    <a id='back' className='text-blue-600 hover:underline hover:underline-offset-1'>Go back</a>
                    <h1 className='text-3xl font-bold'>Random Code Generator</h1>
                    <p className='mt-5 font-medium text-lg'>The generated codes will be stored in our database and you are going to be able to use it as a payment method</p>
                    <h1 id='gift_code' className='mt-10 font-bold text-3xl'></h1>
                    <button
                        onClick={request}
                        className='bg-red-600 rounded mt-10 px-5 py-3 text-white font-bold text-lg hover:bg-red-700 hover:text-black'
                    >
                        Generate
                    </button>
                </div>
            </section>

            <div>
                <FooterSign />
            </div>
        </main>
    )
}

export default GenerateGiftCode;
