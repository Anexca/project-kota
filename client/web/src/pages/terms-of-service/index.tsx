import Container from "../../componnets/shared/container";
import Footer from "../../componnets/shared/footer";
import Header from "../../componnets/shared/header/header";

const TermsOfService = () => {
  return (
    <div>
      <Header />
      <Container className="p-4">
        <h1 className="text-3xl font-bold text-center mb-8 mt-8">
          Terms of Service
        </h1>
        <p className="text-sm text-gray-500 mb-4">Last Updated: 9-Sept-2024</p>

        <p className="mb-6">
          Welcome to <span className="font-semibold">pseudotest.pro</span> (the
          "Service"), an online platform that facilitates online exams for
          students. By accessing or using our website, you agree to comply with
          and be bound by the following terms and conditions. Please review them
          carefully before using the Service.
        </p>

        <h2 className="text-xl font-semibold mb-4">1. Acceptance of Terms</h2>
        <p className="mb-6">
          By using <span className="font-semibold">pseudotest.pro</span>, you
          agree to these Terms of Service. If you do not agree with any of these
          terms, please do not use our Service.
        </p>

        <h2 className="text-xl font-semibold mb-4">2. Changes to the Terms</h2>
        <p className="mb-6">
          We reserve the right to update or modify these Terms at any time
          without prior notice. Changes will be effective immediately upon
          posting. Your continued use of the Service after any modifications
          indicates your acceptance of the new Terms.
        </p>

        <h2 className="text-xl font-semibold mb-4">3. Eligibility</h2>
        <p className="mb-6">
          You must be a registered user and meet the minimum age or educational
          requirements to take exams on our platform. By using the Service, you
          confirm that you meet these requirements.
        </p>

        <h2 className="text-xl font-semibold mb-4">4. Use of the Service</h2>
        <ul className="list-disc list-inside mb-6">
          <li>
            The Service is intended solely for conducting MCQ exams. You may use
            the Service only for the purpose of taking exams.
          </li>
          <li>
            You agree not to use the Service for any illegal or unauthorized
            purposes.
          </li>
          <li>
            You are responsible for your actions and the integrity of your
            responses during the exams.
          </li>
        </ul>

        <h2 className="text-xl font-semibold mb-4">5. User Accounts</h2>
        <ul className="list-disc list-inside mb-6">
          <li>
            You may be required to create an account to take exams. It is your
            responsibility to provide accurate information during registration.
          </li>
          <li>
            You must safeguard your account information and not share it with
            others. You are responsible for any activities conducted under your
            account.
          </li>
        </ul>

        <h2 className="text-xl font-semibold mb-4">6. Exam Conduct</h2>
        <ul className="list-disc list-inside mb-6">
          <li>
            By using our Service, you agree to follow the rules and instructions
            provided for each exam.
          </li>
          <li>
            Cheating, plagiarism, or any form of dishonesty during exams is
            strictly prohibited.
          </li>
          <li>
            We reserve the right to monitor and review exam submissions for any
            suspicious or fraudulent activity.
          </li>
        </ul>

        <h2 className="text-xl font-semibold mb-4">7. Intellectual Property</h2>
        <p className="mb-6">
          All content on the website, including exam questions, results, and
          materials, is the intellectual property of{" "}
          <span className="font-semibold">pseudotest.pro</span>. You may not
          copy, distribute, or share any exam content without our express
          written permission.
        </p>

        <h2 className="text-xl font-semibold mb-4">
          8. Disclaimer of Warranties
        </h2>
        <p className="mb-6">
          The Service is provided "as is" without any warranties, express or
          implied. We make no representations or warranties regarding the
          accuracy, availability, or reliability of the exam content or the
          results generated.
        </p>

        <h2 className="text-xl font-semibold mb-4">
          9. Limitation of Liability
        </h2>
        <p className="mb-6">
          In no event shall{" "}
          <span className="font-semibold">pseudotest.pro</span> be liable for
          any direct, indirect, incidental, special, or consequential damages
          arising from your use of the Service or the results of any exam.
        </p>

        <h2 className="text-xl font-semibold mb-4">10. Termination</h2>
        <p className="mb-6">
          We reserve the right to terminate or suspend your access to the
          Service if you violate these Terms or engage in any prohibited
          activities.
        </p>

        <h2 className="text-xl font-semibold mb-4">11. Governing Law</h2>
        <p className="mb-6">
          These Terms shall be governed by and construed in accordance with the
          laws of India, without regard to its conflict of law principles.
        </p>

        <h2 className="text-xl font-semibold mb-4">12. Contact Information</h2>
        <p className="mb-6">
          If you have any questions about these Terms of Service, please contact
          Harshal Dharmik or Rahil Sheikh at{" "}
          <span className="font-semibold">support@pseudotest.pro</span>.
        </p>
      </Container>
      <Footer />
    </div>
  );
};

export default TermsOfService;
