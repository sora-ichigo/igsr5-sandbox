import React from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/css";

export const Demo: React.FC = () => {
  return (
    <>
      <Swiper
        slidesPerView={"auto"}
        spaceBetween={10}
        slidesPerGroup={1}
        loop={true}
        loopFillGroupWithBlank={true}
        centeredSlides={true}
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
