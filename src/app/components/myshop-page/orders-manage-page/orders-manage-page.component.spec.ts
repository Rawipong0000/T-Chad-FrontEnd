import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrdersManagePageComponent } from './orders-manage-page.component';

describe('OrdersManagePageComponent', () => {
  let component: OrdersManagePageComponent;
  let fixture: ComponentFixture<OrdersManagePageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [OrdersManagePageComponent]
    });
    fixture = TestBed.createComponent(OrdersManagePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
