import Container from "../../componnets/shared/container";
import Footer from "../../componnets/shared/footer";
import Header from "../../componnets/shared/header/header";

const ContactUs = () => {
  return (
    <div className="bg-slate-50 min-h-screen">
      <Header />
      <Container>
        <div className="mt-8">
          <div className="bg-white shadow rounded-lg p-8 text-center">
            <h1 className="text-2xl font-bold text-center mb-8">Contact Us</h1>

            <p className="text mb-8 text-center">
              If you have any questions, issues, or feedback, feel free to reach
              out to us at the email address below. We're here to help!
            </p>

            <p className="text-md font-semibold mb-4">Email Us At:</p>
            <a
              href="mailto:support@pseudotest.pro"
              className="text-info text-md font-bold hover:text-indigo-800"
            >
              support@pseudotest.pro
            </a>
          </div>
        </div>
      </Container>
      <Footer />
    </div>
  );
};

export default ContactUs;
