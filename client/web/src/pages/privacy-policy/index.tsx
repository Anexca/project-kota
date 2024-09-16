import Container from "../../componnets/shared/container";
import Footer from "../../componnets/shared/footer";
import Header from "../../componnets/shared/header/header";

const PrivacyPolicy = () => {
  return (
    <div>
      <Header />
      <Container>
        <div className="p-4">
          <h1 className="text-3xl font-bold text-center mb-8 mt-8">
            Privacy Policy
          </h1>
          <p className="text-sm text-gray-500 mb-4">Last Updated: [Date]</p>

          <p className="mb-6">
            At <span className="font-semibold">pseudotest.pro</span> (the
            "Service"), we are committed to protecting your privacy. This
            Privacy Policy explains how we collect, use, disclose, and safeguard
            your information when you use our website and services. Please read
            this policy carefully.
          </p>

          <h2 className="text-xl font-semibold mb-4">
            1. Information We Collect
          </h2>
          <p className="mb-6">
            We may collect personal information such as your name, email
            address, and educational details when you register on our platform
            to take exams. We also collect non-personal information such as your
            browser type, IP address, and usage patterns.
          </p>

          <h2 className="text-xl font-semibold mb-4">
            2. How We Use Your Information
          </h2>
          <ul className="list-disc list-inside mb-6">
            <li>
              To facilitate your access to exams and other services on our
              platform.
            </li>
            <li>
              To communicate with you regarding your account, exam results, and
              updates to our services.
            </li>
            <li>
              To improve our platform by analyzing user behavior and feedback.
            </li>
            <li>
              To comply with legal requirements and protect the security of the
              platform.
            </li>
          </ul>

          <h2 className="text-xl font-semibold mb-4">
            3. Sharing of Information
          </h2>
          <p className="mb-6">
            We do not share your personal information with third parties except
            in the following cases:
          </p>
          <ul className="list-disc list-inside mb-6">
            <li>When required by law or legal process.</li>
            <li>
              To service providers who assist us in operating the platform,
              provided they agree to keep your information confidential.
            </li>
            <li>
              In the event of a business transfer, such as a merger or
              acquisition, your information may be transferred as part of the
              assets.
            </li>
          </ul>

          <h2 className="text-xl font-semibold mb-4">4. Data Security</h2>
          <p className="mb-6">
            We use administrative, technical, and physical security measures to
            protect your personal information. However, please note that no
            system is 100% secure, and we cannot guarantee the absolute security
            of your information.
          </p>

          <h2 className="text-xl font-semibold mb-4">
            5. Cookies and Tracking Technologies
          </h2>
          <p className="mb-6">
            We may use cookies and similar tracking technologies to enhance your
            experience on our website, monitor user activity, and analyze usage
            patterns. You can manage your cookie preferences through your
            browser settings.
          </p>

          <h2 className="text-xl font-semibold mb-4">6. Your Privacy Rights</h2>
          <p className="mb-6">
            You have the right to access, update, or delete your personal
            information at any time. If you wish to exercise any of these
            rights, please contact Harshal Dharmik or Rahil Sheikh at
            support@pseudotest.pro.
          </p>

          <h2 className="text-xl font-semibold mb-4">7. Children's Privacy</h2>
          <p className="mb-6">
            Our Service is not intended for children under the age of 13. We do
            not knowingly collect personal information from children under 13.
            If you are a parent or guardian and believe we have collected such
            information, please contact Harshal Dharmik or Rahil Sheikh at
            support@pseudotest.pro to request its deletion.
          </p>

          <h2 className="text-xl font-semibold mb-4">
            8. Changes to This Policy
          </h2>
          <p className="mb-6">
            We may update this Privacy Policy from time to time. Any changes
            will be posted on this page with a revised "Last Updated" date. Your
            continued use of the Service after such changes constitutes your
            acceptance of the new policy.
          </p>

          <h2 className="text-xl font-semibold mb-4">9. Contact Us</h2>
          <p className="mb-6">
            If you have any questions about this Privacy Policy or your personal
            information, please contact Harshal Dharmik or Rahil Sheikh at{" "}
            <span className="font-semibold">support@pseudotest.pro</span>.
          </p>
        </div>
      </Container>
      <Footer />
    </div>
  );
};

export default PrivacyPolicy;
