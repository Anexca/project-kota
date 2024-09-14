import { Link } from "react-router-dom";
import Icon from "../../base/icon";
import Container from "../container";
import Footer from "../footer";
import Header from "../header/header";
import { AnimatedBeamHeading } from "./beam-heading";
import { paths } from "../../../routes/route.constant";

const UserHomePage = () => {
  return (
    <>
      <Header />
      <div className="relative" id="home">
        <div
          aria-hidden="true"
          className="absolute inset-0 grid grid-cols-2 -space-x-52 opacity-40 dark:opacity-20"
        >
          <div className="blur-[106px] h-56 bg-gradient-to-br from-primary to-purple-400 dark:from-blue-700"></div>
          <div className="blur-[106px] h-32 bg-gradient-to-r from-cyan-400 to-sky-300 dark:to-indigo-600"></div>
        </div>
        <div className="max-w-7xl mx-auto px-6 md:px-12 xl:px-6">
          <div className="relative ml-auto pt-20">
            <div className="lg:w-2/3 text-center mx-auto min-h-[80vh]">
              <div className="text-lg mb-2">Your Go-To Place for </div>
              <div className="text-2xl text-info font-semibold mb-4">
                Complete Exam Preparation
              </div>
              <div className="flex items-center justify-center mt-[10%]">
                <AnimatedBeamHeading />
              </div>

              <div className="mt-20 font-semibold">Take a test now.</div>
              <div>
                <Icon icon="arrow_right" className="text-info rotate-90" />
              </div>
            </div>
            <div className="w-full bg-white">
              <Container>
                <h4 className="border-b py-2 font-bold text-center">
                  Available Test
                </h4>
                <div className="p-2 md:p-4">
                  <Link to={`/${paths.EXAMS}/banking/${paths.DISCRIPTIVE}`}>
                    <div className="shadow rounded-sm p-4 border">
                      <div className="font-semibold text-base">
                        Banking Descriptive Exams
                      </div>
                      <div className=" text-sm">
                        AI-Assisted Descriptive Assessment with Insights. Get
                        understanding about your weak and strong points.
                      </div>
                    </div>
                  </Link>
                </div>
              </Container>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default UserHomePage;
