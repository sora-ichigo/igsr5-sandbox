import React from "react";
import "swiper/css";
import "swiper/css/navigation";
import SwiperCore, { Pagination, Navigation } from "swiper";

SwiperCore.use([Pagination, Navigation]);

const App = ({ Component, pageProps }: any) => {
  return <Component {...pageProps} />;
};

export default App;
