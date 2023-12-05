import { google } from "googleapis";

export default function Page() {
  async function create(formData: FormData) {
    "use server";

    const OAuth2 = google.auth.OAuth2;
    const oauth2Client = new OAuth2(
      process.env.CLIENT_ID,
      process.env.CLIENT_SECRET,
      process.env.REDIRECT_URL
    );

    await fetch("https://api.resend.com/emails", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${process.env.RESEND_KEY}`,
      },
      body: JSON.stringify({
        from: "Someone <dearselfapp@gmail.com>",
        to: formData.get("to"),
        subject: formData.get("subject"),
        text: formData.get("text"),
      }),
    });
    console.log("Email sent!");
  }
  return (
    <div className="dark:bg-gray-800 flex items-center justify-center min-h-screen py-12 px-4 sm:px-6 lg:px-8">
    <form action={create} className="mt-8 space-y-6 w-full max-w-md">
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
        <input type="date" name="scheduleDate" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-4 py-2" />
      </label>
      <label className="block text-sm font-medium text-gray-700 dark:text-gray-200">
        Schedule Time
        <input type="time" name="scheduleTime" className="mt-1 block w-full dark:bg-gray-700 dark:text-white dark:border-gray-600 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-4 py-2" />
      </label>
      <button type="submit" className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Submit</button>
    </form>
  </div>
  );
}