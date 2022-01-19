import React, { useCallback, useEffect, useRef, useState } from "react";
import Link from "next/link";
import styled from "styled-components";
import type CoreSwiper from "swiper";
import { Swiper, SwiperSlide, SwiperProps } from "swiper/react";

const pages = ["1", "2", "3", "4", "5"];

type CustomSwiperSlideProps = {
  isActive: boolean;
  isNext: boolean;
  isPrev: boolean;
};

export const Demo: React.FC = () => {
  // NOTE: useSwiper みたいな形で関連処理まとめられそう
  const [swiper, setSwiper] = useState<CoreSwiper | undefined>(undefined);
  const [tabSwiper, setTabSwiper] = useState<CoreSwiper | undefined>(undefined);

  const slideStatus = useRef<{ prev: number; current: number }>({
    prev: 0,
    current: 0,
  });

  const swiperGeneralProps: SwiperProps = {
    slidesPerView: "auto",
    spaceBetween: 10,
    loop: true,
    speed: 100,
    centeredSlides: true,
  };

  const onMainSlideChange = (swiper: CoreSwiper) => {
    // TODO: 超高速スワイプすると tab の変更が追いつかない
    // 現在は slide speed を上げて追いつけるようにしているけど、見た目があんまりよろしくない
    // pc サイズの時はスワイプ切るのもあり
    slideStatus.current = {
      prev: slideStatus.current.current,
      current: swiper.realIndex,
    };
    const op = checkSlideMove();
    tabSlideMove(op);
  };

  // TODO: テスト書く
  const checkSlideMove = useCallback<() => "NOOP" | "NEXT" | "PREV">(() => {
    if (slideStatus.current.prev === slideStatus.current.current) {
      // noop
      return "NOOP";
    }

    if (
      slideStatus.current.current === 0 &&
      slideStatus.current.prev === pages.length - 1
    ) {
      return "NEXT";
    }

    if (
      slideStatus.current.current === pages.length - 1 &&
      slideStatus.current.prev === 0
    ) {
      return "PREV";
    }

    if (slideStatus.current.current > slideStatus.current.prev) {
      return "NEXT";
    } else {
      return "PREV";
    }
  }, [slideStatus]);

  const tabSlideMove = (op: "NOOP" | "NEXT" | "PREV") => {
    switch (op) {
      case "NEXT":
        tabSwiper?.slideNext();
        break;
      case "PREV":
        tabSwiper?.slidePrev();
        break;
      default:
      // noop
    }
  };

  const slidePrev = () => swiper?.slidePrev();
  const slideNext = () => swiper?.slideNext();

  const AttachOverray: React.FC<CustomSwiperSlideProps> = ({
    isNext,
    isPrev,
  }) => (
    <>
      {isNext ? (
        <OverrayDiv onClick={() => slideNext()} />
      ) : isPrev ? (
        <OverrayDiv onClick={() => slidePrev()} />
      ) : null}
    </>
  );

  return (
    <>
      <Swiper
        {...swiperGeneralProps}
        onInit={(swiper: CoreSwiper) => setTabSwiper(swiper)}
        allowTouchMove={false}
        className="tab"
      >
        {pages.map((v) => (
          <SwiperSlide style={{ width: "30%" }} key={v}>
            <div style={{ textAlign: "center" }}>
              <Link href="/">
                <a>tab{v}</a>
              </Link>
            </div>
          </SwiperSlide>
        ))}
      </Swiper>

      <div style={{ height: "50px" }}></div>

      <Swiper
        {...swiperGeneralProps}
        onInit={(swiper: CoreSwiper) => setSwiper(swiper)}
        onSlideChange={onMainSlideChange}
        className="mainSlide"
      >
        {pages.map((v) => (
          <SwiperSlide style={{ width: "80%" }} key={v}>
            {(props) => (
              <StyledDiv>
                <AttachOverray {...props} />
                <Link href="/">
                  <a>slide{v}</a>
                </Link>
              </StyledDiv>
            )}
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

const StyledDiv = styled.div`
  height: 500px;
  text-align: center;
`;

const OverrayDiv = styled.div`
  position: absolute;
  width: 100%;
  height: 100%;
  cursor: pointer;
  background-color: #ffffff9d;
  z-index: 99999;
`;
