import { Component, OnInit } from '@angular/core';
import { MyShopService } from 'src/app/service/myshop.service';
import { Router } from '@angular/router';
import { Ordering } from 'src/app/model/transaction.model';
import { KeyValue } from '@angular/common';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-orders-manage-page',
  templateUrl: './orders-manage-page.component.html',
  styleUrls: ['./orders-manage-page.component.css']
})
export class OrdersManagePageComponent implements OnInit {

  orders: Ordering[] = [];
  tracking: { [tranKey: string]: string } = {};
  public objectKeys = Object.keys;
  nestedGroupedOrders: Record<string, Record<string, Ordering[]>> = {};

  constructor(
    private myShopService: MyShopService,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.myShopService.getMyShopTransaction().subscribe({
      next: (data) => {
        console.log('Fetched orders data:', data);
        this.orders = data.map((order: Ordering) => ({
          ...order,
          create_date: new Date(order.create_date),
          update_date: new Date(order.update_date),
        }));
        this.GroupingOrder();
      },
      error: (err) => {
        console.error('Failed to fetch orders', err);
      }
    });
  }

  GroupingOrder() {
    this.nestedGroupedOrders = this.orders.reduce((acc, item) => {
      const dateKey = item.create_date.toISOString().split('T')[0]; // เช่น "2025-07-27"
      const tranKey = item.sub_tran_id.toString();

      // สร้างกลุ่มวันที่ถ้ายังไม่มี
      if (!acc[dateKey]) {
        acc[dateKey] = {};
      }

      // สร้างกลุ่ม tran_id ภายใต้วันที่ ถ้ายังไม่มี
      if (!acc[dateKey][tranKey]) {
        acc[dateKey][tranKey] = [];
      }

      acc[dateKey][tranKey].push(item);

      return acc;
    }, {} as Record<string, Record<string, Ordering[]>>);
  }

  EditTrackingNo(tranKey: string) {
    const track = this.tracking[tranKey] || '';

    if (!track) {
      alert('tracking cannot be blank');
    } else {
      const body = {
        Tran_id: +tranKey,
        Tracking: track
      };

      console.log('Tracking:', body);

      this.myShopService.editTracking(body).subscribe({
        next: (response) => {
          console.log('Update successful:', response);
          alert('Update successful!');
          this.ngOnInit()
        },
        error: (err) => {
          console.error("update failed:", err.error);
          alert("update failed: " + err.error);
        }
      })
    }
  }

  ApproveRefund(tranKey: string){
    const SubTranID = +tranKey;
    this.myShopService.approveRefund(SubTranID).subscribe({
      next: (response) => {
        console.log('Refund approved:', response);
        alert('Refund approved!');
        this.ngOnInit()
      },
      error: (err) => {
        console.error("refund failed:", err.error);
        alert("refund failed: " + err.error);
      }
    })
  }

  RejectRefund(tranKey: string){
    const SubTranID = +tranKey;
    this.myShopService.rejectRefund(SubTranID).subscribe({
      next: (response) => {
        console.log('Reject approved:', response);
        alert('Reject approved!');
        this.ngOnInit()
      },
      error: (err) => {
        console.error("reject failed:", err.error);
        alert("reject failed: " + err.error);
      }
    })
  }

  CancelTransaction(tranKey: string){
    const SubTranID = +tranKey;
    this.myShopService.cancelTransaction(SubTranID).subscribe({
      next: (response) => {
        console.log('Cancel successful:', response);
        alert('Cancel successful!');
        this.ngOnInit()
      },
      error: (err) => {
        console.error("cancel failed:", err.error);
        alert("cancel failed: " + err.error);
      }
    })
  }

  goToOrderManage(){
    this.router.navigate(['/MyShop/user/orders-manage-page']);
  }

  dateDesc = (a: KeyValue<string, any>, b: KeyValue<string, any>): number => {
      return new Date(b.key).getTime() - new Date(a.key).getTime();
    };
  
    tranDesc = (a: KeyValue<string, any>, b: KeyValue<string, any>): number => {
      return Number(b.key) - Number(a.key);
    };
}


