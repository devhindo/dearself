"use client";

import { create } from "./send";

export default function ClientButton() {
    return (
        <div>
            <div className="dark:bg-gray-800 flex flex-col items-center justify-center min-h-screen py-12 px-4 sm:px-6 lg:px-8">
                <h1 className="text-4xl font-bold text-indigo-600 mb-8">dearself &#x1F60E;</h1>
                <header className="text-center mb-8">
                    <p className="text-lg text-gray-700 dark:text-gray-200">
                        This website allows you to send future emails to yourself or others.
                    </p>
                </header>
                <form action={create} className="mt-8 space-y-6 w-full max-w-md">
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-200">
                        Name
                        <input type="text" name="name" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-4 py-2" />
                    </label>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-200">
                        Subject
                        <input type="text" name="subject" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-4 py-2" />
                    </label>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-200">
                        To
                        <input type="text" name="to" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-4 py-2" />
                    </label>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-200">
                        Text
                        <textarea name="text" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm w-full h-40 px-4 py-2" />
                    </label>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-200">
                        Schedule Date
                        <input type="date" name="date" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-4 py-2" />
                    </label>
                    <button
                        className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        onClick={async () => {
                            const formData = new FormData(document.querySelector("form")!);
                            const response = await create(formData);
                            console.log("resss")
                            console.log("resbponse" + JSON.stringify(response))
                            if (response.error) {
                                alert(response.error);
                            } else {
                                alert(response.message);
                            }
                        }}
                    >
                        Send Mail
                    </button>
                </form>
                <footer className="mt-8 text-gray-500 dark:text-gray-400 text-center">
                    Made by <a href="https://github.com/devhindo" target="_blank" rel="noopener noreferrer" className="underline">devhindo</a> on <a href="https://github.com/devhindo/dearself" target="_blank" rel="noopener noreferrer" className="underline">GitHub</a>
                </footer>
            </div>
        </div>
    );
}