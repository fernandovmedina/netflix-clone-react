import FooterSign from '../../components/FooterSign.tsx'
import NavbarSign from '../../components/NavbarSign.tsx'

const GenerateGiftCode = () => {
    
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
            console.error('Error al hacer la solicitud:', error);
        }
    }

    return (
        <main>
            <div>
                <NavbarSign />
            </div>

            <section className='w-full flex items-center justify-center my-20'>
                <div className='flex flex-col items-center'>
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
