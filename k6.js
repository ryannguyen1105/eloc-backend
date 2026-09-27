import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '5s', target: 10 },   // Khởi động nhẹ 10 VUs
    { duration: '20s', target: 30 },  // Duy trì 30 VUs liên tục tạo sản phẩm
    { duration: '5s', target: 0 },    // Hạ nhiệt về 0
  ],
  thresholds: {
    'http_req_failed': ['rate<0.01'],   // Tỉ lệ lỗi dưới 1%
    'http_req_duration': ['p(95)<300'], // 95% request hoàn thành dưới 300ms
  },
};

const BASE_URL = 'http://localhost:8081';

export default function () {
  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const uniqueId = `${Date.now()}_${__VU}_${__ITER}`;
  
  const payload = JSON.stringify({
    category_id: 1, // Đảm bảo ID danh mục này đã tồn tại trong DB
    name: `Sản Phẩm ${uniqueId}`,
    slug: `san-pham-${uniqueId}`,
    sku: `SKU-${uniqueId}`,
    price: Math.floor(Math.random() * 500000) + 50000,
    stock: 100,
    description: `Mô tả chi tiết cho sản phẩm ${uniqueId}. Chất lượng cao, chính hãng.`,
    status: 'active', // Thêm trường status tương ứng với DB
    attributes: {
      color: "Test",
      size: "Test",
      material: "Test"
    }
  });

  const res = http.post(`${BASE_URL}/product`, payload, params);

  if (res.status !== 200 && res.status !== 201) {
    console.log(`[Create Product Failed]: ${res.status} - ${res.body}`);
  }

  check(res, {
    'create product status 200/201': (r) => r.status === 200 || r.status === 201,
  });

  sleep(0.5);
}