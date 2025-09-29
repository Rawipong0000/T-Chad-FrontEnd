import { Component } from '@angular/core';
import { HistoryService } from 'src/app/service/history.service';
import { Router } from '@angular/router';
import { HistoryOrdering } from 'src/app/model/transaction.model';
import { KeyValue } from '@angular/common';

@Component({
  selector: 'app-history-page',
  templateUrl: './history-page.component.html',
  styleUrls: ['./history-page.component.css']
})
export class HistoryPageComponent {

  orders: HistoryOrdering[] = [];
  tracking: { [tranKey: string]: string } = {};
  public objectKeys = Object.keys;
  nestedGroupedOrders: Record<string, Record<string, HistoryOrdering[]>> = {};

  constructor(
    private historyService: HistoryService,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.historyService.getHistoryTransaction().subscribe({
      next: (data) => {
        console.log('Fetched orders data:', data);
        this.orders = data.map((order: HistoryOrdering) => ({
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
    }, {} as Record<string, Record<string, HistoryOrdering[]>>);
  }

  CompleteTransaction(tranKey: string) {
    const SubTranID = +tranKey;
    this.historyService.completeTransaction(SubTranID).subscribe({
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

  RefundTransaction(tranKey: string) {
    const SubTranID = +tranKey;
    this.historyService.refundTransaction(SubTranID).subscribe({
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

  goToOrderManage() {
    this.router.navigate(['/MyShop/user/orders-manage-page']);
  }

  dateDesc = (a: KeyValue<string, any>, b: KeyValue<string, any>): number => {
    return new Date(b.key).getTime() - new Date(a.key).getTime();
  };

  tranDesc = (a: KeyValue<string, any>, b: KeyValue<string, any>): number => {
    return Number(b.key) - Number(a.key);
  };
}
