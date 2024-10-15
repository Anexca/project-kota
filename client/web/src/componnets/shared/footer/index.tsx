import { Link } from "react-router-dom";
import { paths } from "../../../routes/route.constant";

const Footer = () => {
  return (
    <footer className="bg-white rounded-lg  p-4 dark:bg-gray-800 border-t border-t-gray-100">
      <div className="w-full mx-auto max-w-screen-xl p-4 md:flex md:items-center md:justify-between">
        <span className="text-sm text-gray-500 sm:text-center dark:text-gray-400">
          © 2024{" "}
          <Link to="https://pseudotest.pro/" className="hover:underline">
            PseudoTest
          </Link>
          . All Rights Reserved.
        </span>
        <ul className="flex flex-wrap items-center mt-3 text-sm font-medium text-gray-500 dark:text-gray-400 sm:mt-0">
          <li>
            <Link
              to={`/${paths.CONTACT_US}`}
              className="hover:underline me-4 md:me-6"
            >
              Contact Us
            </Link>
          </li>
          <li>
            <Link
              to={`/${paths.TERMS_OF_SERVICE}`}
              className="hover:underline me-4 md:me-6"
            >
              Terms Of Service
            </Link>
          </li>
          <li>
            <Link
              to={`/${paths.PRIVACY_POLICY}`}
              className="hover:underline me-4 md:me-6"
            >
              Privacy Policy
            </Link>
          </li>
        </ul>
      </div>
    </footer>
  );
};

export default Footer;
