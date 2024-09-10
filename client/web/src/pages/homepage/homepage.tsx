import { Link } from "react-router-dom";
import Header from "../../componnets/shared/header/header";
import { paths } from "../../routes/route.constant";
import Footer from "../../componnets/shared/footer";

const card = (
  <article className="relative block overflow-hidden rounded-lg border border-gray-100 p-4 sm:p-6 lg:p-8">
    <span className="absolute inset-x-0 bottom-0 h-2 bg-gradient-to-r from-green-300 via-blue-500 to-purple-600"></span>

    <div className="sm:flex sm:justify-between sm:gap-4">
      <div>
        <h3 className="text-lg font-bold text-gray-900 sm:text-xl">
          Building a SaaS product as a software developer
        </h3>

        <p className="mt-1 text-xs font-medium text-gray-600">By John Doe</p>
      </div>

      <div className="hidden sm:block sm:shrink-0">
        <img
          alt=""
          src="https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1180&q=80"
          className="size-16 rounded-lg object-cover shadow-sm"
        />
      </div>
    </div>

    <div className="mt-4">
      <p className="text-pretty text-sm text-gray-500">
        Lorem ipsum dolor sit, amet consectetur adipisicing elit. At velit illum
        provident a, ipsa maiores deleniti consectetur nobis et eaque.
      </p>
    </div>

    <dl className="mt-6 flex gap-4 sm:gap-6">
      <div className="flex flex-col-reverse">
        <dt className="text-sm font-medium text-gray-600">Published</dt>
        <dd className="text-xs text-gray-500">31st June, 2021</dd>
      </div>

      <div className="flex flex-col-reverse">
        <dt className="text-sm font-medium text-gray-600">Reading time</dt>
        <dd className="text-xs text-gray-500">3 minute</dd>
      </div>
    </dl>
  </article>
);

const HomePage = () => {
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
            <div className="lg:w-2/3 text-center mx-auto ">
              <h1 className="text-balance text-gray-900 dark:text-white font-bold text-5xl md:text-6xl xl:text-7xl">
                Ace Your Exams with{" "}
                <span className="text-purple-500 dark:text-white">
                  AI-Powered Mock
                </span>{" "}
                Tests
              </h1>
              <p className="mt-8 text-gray-700 dark:text-gray-300">
                Experience the next generation of mock tests designed to give
                you an edge in your exam preparation. Our platform offers
                in-depth practice for various goverment and entrance exam with
                an exclusive feature of AI-evaluated descriptive answers.
              </p>
              <div className="mt-16 flex flex-wrap justify-center gap-y-4 gap-x-6">
                <Link
                  to={`/${paths.REGISTER}`}
                  className="relative flex h-11 w-full items-center justify-center px-6 before:absolute before:inset-0 before:rounded-full before:bg-primary before:transition before:duration-300 hover:before:scale-105 active:duration-75 active:before:scale-95 sm:w-max"
                >
                  <span className="relative text-base font-semibold text-white">
                    Register to Get started
                  </span>
                </Link>
                <Link
                  to={`/${paths.LOGIN}`}
                  className="relative flex h-11 w-full items-center justify-center px-6 before:absolute before:inset-0 before:rounded-full before:border before:border-transparent before:bg-primary/10 before:bg-gradient-to-b before:transition before:duration-300 hover:before:scale-105 active:duration-75 active:before:scale-95 dark:before:border-gray-700 dark:before:bg-gray-800 sm:w-max"
                >
                  <span className="relative text-base font-semibold text-primary dark:text-white">
                    Log In
                  </span>
                </Link>
              </div>
              <div className="hidden py-8 mt-16 border-y border-gray-100 dark:border-gray-800 sm:flex justify-between  gap-2">
                <div className="text-left flex-1">
                  <h6 className="text-lg font-semibold text-gray-700 dark:text-white">
                    Lowest price
                  </h6>
                  <p className="mt-2 text-gray-500">
                    Get top-tier mock tests without breaking the bank.
                  </p>
                </div>
                <div className="text-left flex-1">
                  <h6 className="text-lg font-semibold text-gray-700 dark:text-white">
                    Fastest on the market
                  </h6>
                  <p className="mt-2 text-gray-500">
                    Experience lightning-fast results and feedback.
                  </p>
                </div>
                <div className="text-left flex-1">
                  <h6 className="text-lg font-semibold text-gray-700 dark:text-white">
                    Most loved
                  </h6>
                  <p className="mt-2 text-gray-500">
                    Our platform delivers AI-powered evaluations that students
                    rave about.
                  </p>
                </div>
              </div>
            </div>
            <div className="mt-12 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3">
              <div className="p-4 grayscale transition duration-200 hover:grayscale-0">
                {card}
              </div>
              <div className="p-4 grayscale transition duration-200 hover:grayscale-0">
                {card}
              </div>
              <div className="p-4 grayscale transition duration-200 hover:grayscale-0">
                {card}
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default HomePage;
