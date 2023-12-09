const send_email_api_url = "http://localhost:8080/send"
//const send_email_api_url = "https://dearself.onrender.com/send"

export default function Page() {
  async function create(formData: FormData) {
    "use server";
    console.log(formData);
    const response = await fetch(send_email_api_url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        
      },
      body: JSON.stringify({
        name: formData.get("name"),
        to: formData.get("to"),
        subject: formData.get("subject"),
        text: formData.get("text"),
        date: formData.get("date"),
      }),
    });

    if (response.status === 200) {
      console.log("Email scheduled successfully!")
    } else {
      console.log("There was an error scheduling the mail.")
    }

    console.log("Email sent!");
  }
  return (
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
        <button type="submit" className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">send Email</button>
      </form>
      <footer className="mt-8 text-gray-500 dark:text-gray-400 text-center">
        Made by <a href="https://github.com/devhindo" target="_blank" rel="noopener noreferrer" className="underline">devhindo</a> on <a href="https://github.com/devhindo/dearself" target="_blank" rel="noopener noreferrer" className="underline">GitHub</a>
      </footer>
    </div>
  );
}