import { test, expect } from "bun:test";
import { readTextFile } from "../utils";

type OrganisedPageOrders = Record<number, number[]>;

function sort(line: string, orders: OrganisedPageOrders): string {
  const pages = line.split(",").filter(Boolean);
  let checked = "";
  for (let i = 0; i < pages.length; i++) {
    if (!pages[i + 1]) return (checked += pages[i]);

    const numA = parseInt(pages[i]);
    const numB = parseInt(pages[i + 1]);
    if (orders[numA] && orders[numA].includes(numB)) {
      checked += `${pages[i]},`;
    }
  }
  return checked;
}

function organiseOrders(pageOrders: string) {
  const ordered: OrganisedPageOrders = {};

  pageOrders.split("\n").forEach((order) => {
    const [before, after] = order.split("|");
    if (ordered[+before]) {
      ordered[+before].push(+after);
    } else {
      ordered[+before] = [];
      ordered[+before].push(+after);
    }
  });

  return ordered;
}

test("part 1 example", () => {
  const correctLines: string[] = [];
  const [pageOrders, updates] = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
    .split("\n\n")
    .filter(Boolean);

  const ordered = organiseOrders(pageOrders);

  updates
    .split("\n")
    .filter(Boolean)
    .forEach((line) => {
      const sortedLine = sort(line, ordered);
      line === sortedLine && correctLines.push(line);
    });

  const sum = correctLines.reduce((acc, curr) => {
    const line = curr.split(",");
    const idx = Math.round((line.length - 1) / 2);
    const mid = +line[idx];
    return (acc += mid);
  }, 0);

  expect(sum).toBe(143);
});

test("part 1", async () => {
  const correctLines: string[] = [];
  const [pageOrders, updates] = (await readTextFile("./input.txt"))
    .split("\n\n")
    .filter(Boolean);

  const ordered = organiseOrders(pageOrders);

  updates
    .split("\n")
    .filter(Boolean)
    .forEach((line) => {
      const sortedLine = sort(line, ordered);
      line === sortedLine && correctLines.push(line);
    });

  const sum = correctLines.reduce((acc, curr) => {
    const line = curr.split(",");
    const idx = Math.round((line.length - 1) / 2);
    const mid = +line[idx];
    return (acc += mid);
  }, 0);

  expect(sum).toBe(6260);
});
