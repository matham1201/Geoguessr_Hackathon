import React from 'react';

function Page(){
    return (
        <main className="bg-celeste p-5 min-h-screen">
            <div className="inline-block max-w-sm mx-auto bg-blue rounded overflow-hidden shadow-lg">
                <div className="px-6 py-20">
                    <div className="font-bold border-white text-xl mb-2"> contenue de la carte </div>
                </div>
                {/* <div class="px-6 py-4">
                    <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2">#Tag1</span>
                    <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2">#Tag2</span>
                    <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700">#Tag3</span>
                </div> */}
            </div>
        </main>
    )
}

export default Page;
