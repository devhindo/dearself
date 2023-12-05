export default function Page() {
  async function create(formData: FormData) {
    "use server";
    await fetch("https://api.resend.com/emails", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${process.env.RESEND_KEY}`,
      },
      body: JSON.stringify({
        from: "Someone <configured@email.provider>",
        to: "contact@launchfa.st",
        subject: formData.get("subject"),
        text: "This works!",
      }),
    });
    console.log("Email sent!");
  }
  return (
    <form action={create}>
      <input type="text" name="subject" />
      <button type="submit">Submit</button>
    </form>
  );
}