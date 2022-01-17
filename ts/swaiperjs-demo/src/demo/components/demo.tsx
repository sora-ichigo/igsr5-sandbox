import React, { useEffect, useState } from "react";
import Link from "next/link";
import styled from "styled-components";
import type CoreSwiper from "swiper";
import { Swiper, SwiperSlide } from "swiper/react";

const pages = ["1", "2", "3", "4", "5"];

type CustomSwiperSlideProps = {
  isActive: boolean;
  isNext: boolean;
  isPrev: boolean;
};

export const Demo: React.FC = () => {
  const [swiper, setSwiper] = useState<CoreSwiper | undefined>(undefined);

  const slidePrev = () => swiper?.slidePrev();
  const slideNext = () => swiper?.slideNext();

  const CustomSwiperSlide: React.FC<CustomSwiperSlideProps> = ({
    children,
    isActive,
    isNext,
    isPrev,
  }) => (
    // TODO: click 連打でバグるから一定間隔イベントを無効化する
    // 自分のポジションを isNext, isPrev で取得しないと言う手もある
    <StyledDiv>
      {isActive ? (
        <>{children}</>
      ) : isNext ? (
        <>
          <OverrayDiv onClick={() => slideNext()} />
          {children}
        </>
      ) : isPrev ? (
        <>
          <OverrayDiv onClick={() => slidePrev()} />
          {children}
        </>
      ) : null}
    </StyledDiv>
  );

  return (
    <>
      <Swiper
        slidesPerView={"auto"}
        spaceBetween={10}
        loop={true}
        centeredSlides={true}
        onInit={(swiper) => setSwiper(swiper)}
        onSlideChange={(swiper) => {
          console.log("slide change. to", swiper.activeIndex);
        }}
        className="mySwiper"
      >
        <SwiperSlide>
          <></>
        </SwiperSlide>
      </Swiper>

      <Swiper
        slidesPerView={"auto"}
        spaceBetween={10}
        loop={true}
        centeredSlides={true}
        onInit={(swiper) => setSwiper(swiper)}
        onSlideChange={(swiper) => {
          console.log("slide change. to", swiper.activeIndex);
        }}
        className="mySwiper"
      >
        {pages.map((v) => (
          <SwiperSlide style={{ width: "80%" }} key={v}>
            {(props) => (
              <CustomSwiperSlide {...props}>
                <Link href="/">
                  <a>slide{v}</a>
                </Link>
              </CustomSwiperSlide>
            )}
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

const StyledDiv = styled.div`
  height: 500px;
`;

const OverrayDiv = styled.div`
  position: absolute;
  width: 100%;
  height: 100%;
  cursor: pointer;
  background-color: #ffffff9d;
  z-index: 99999;
`;
