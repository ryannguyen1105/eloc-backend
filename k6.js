import http from 'k6/http';
import { check, group, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 200 }, // Kéo lên 200 VUs săn sale
    { duration: '20s', target: 500 }, // Đỉnh điểm Flash Sale 500 VUs cùng spam nút Mua
    { duration: '10s', target: 0 },   // Hết giờ sale
  ],
  thresholds: {
    'http_req_failed': ['rate<0.1'],     // Cho phép rớt 10% do giành giật (thực tế Shopee cũng vậy)
    'http_req_duration': ['p(95)<1500'], // Cho phép trễ lên tới 1.5s
  },
};

const BASE_URL = 'http://localhost:8081';

// 🎯 Hàng Flash Sale: TẤT CẢ mọi người đều chọc vào đúng 1 ID này
const FLASH_SALE_PRODUCT_ID = 1; 

export default function () {
  // Mỗi VU là một user khác nhau để tránh lỗi Unique Cart/Order per User
  const userId = (__VU % 1000) + 1; 

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  // 🛒 BƯỚC 1: Đưa hàng Flash Sale vào Giỏ (Spam 500 VUs chọc vào product 1)
  group('1. Add Flash Sale Item to Cart', function () {
    const cartPayload = JSON.stringify({
      user_id: userId,
      product_id: FLASH_SALE_PRODUCT_ID,
      quantity: 1, // Săn sale mỗi người chỉ mua 1 cái
    });

    const resCart = http.post(`${BASE_URL}/cart`, cartPayload, params);
    
    // Nếu rớt, in log để biết lý do DB từ chối
    if (resCart.status !== 200 && resCart.status !== 201) {
      console.log(`[Cart Failed] VU ${__VU}: ${resCart.status} - ${resCart.body}`);
    }

    check(resCart, {
      'add to cart success': (r) => r.status === 200 || r.status === 201,
    });
  });

  // Thời gian user click chuyển màn hình rất nhanh (0.1s - 0.2s)
  sleep(0.1); 

  // 📦 BƯỚC 2: Giành giật Đặt Hàng (Checkout)
  group('2. Flash Sale Checkout', function () {
    const orderPayload = JSON.stringify({
      user_id: userId,
      shipping_address: 'Flash Sale Gò Vấp',
      customer_phone: `090${Math.floor(1000000 + Math.random() * 9000000)}`,
    });

    const resOrder = http.post(`${BASE_URL}/order`, orderPayload, params);

    // Bắt lỗi Deadlock hoặc Out of Stock
    if (resOrder.status !== 200 && resOrder.status !== 201) {
      console.log(`[Order Failed] VU ${__VU}: ${resOrder.status} - ${resOrder.body}`);
    }

    check(resOrder, {
      'checkout success': (r) => r.status === 200 || r.status === 201,
    });
  });

  sleep(0.5); 
}