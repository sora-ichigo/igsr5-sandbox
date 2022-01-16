import React, { useEffect, useState } from "react";
import type CoreSwiper from "swiper";
import { Swiper, SwiperSlide } from "swiper/react";

export const Demo: React.FC = () => {
  const [swiper, setSwiper] = useState<CoreSwiper | undefined>(undefined);

  const slidePrev = () => swiper?.slidePrev();
  const slideNext = () => swiper?.slideNext();

  return (
    <>
      <button onClick={() => slidePrev()}>prev</button>
      <button onClick={() => slideNext()}>next</button>

      <Swiper
        slidesPerView={"auto"}
        spaceBetween={10}
        slidesPerGroup={1}
        loop={true}
        loopFillGroupWithBlank={true}
        centeredSlides={true}
        onInit={(swiper) => setSwiper(swiper)}
        onSlideChange={(swiper) => {
          console.log("slide change. to", swiper.activeIndex);
        }}
        pagination={{
          clickable: true,
        }}
        navigation={true}
        className="mySwiper"
      >
        <SwiperSlide
          style={{ width: "80%", height: "500px", background: "#f0f0f0" }}
        >
          Slide 1
        </SwiperSlide>
        <SwiperSlide
          style={{ width: "80%", height: "500px", background: "#f0f0f0" }}
        >
          Slide 2
        </SwiperSlide>
        <SwiperSlide
          style={{ width: "80%", height: "500px", background: "#f0f0f0" }}
        >
          Slide 3
        </SwiperSlide>
        <SwiperSlide
          style={{ width: "80%", height: "500px", background: "#f0f0f0" }}
        >
          Slide 4
        </SwiperSlide>
      </Swiper>
    </>
  );
};
